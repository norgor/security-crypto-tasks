package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/norgor/security-crypto-tasks/6/2/internal/hashing"
)

type User struct {
	Username     string             `json:"-"`
	PasswordHash hashing.ServerHash `json:"hash"`
}
type Users map[string]*User

var nilUser = User{
	Username:     "",
	PasswordHash: hashing.NewServerHash(""),
}

var errInvalidCredentials = errors.New("user not found or password invalid")

func readUsers() (Users, error) {
	f, err := os.Open("users.json")
	if err != nil {
		return nil, fmt.Errorf("unable to open users file: %w", err)
	}
	defer f.Close()
	users := make(Users)
	if err := json.NewDecoder(f).Decode(&users); err != nil {
		return nil, fmt.Errorf("unable to read users file: %w", err)
	}
	for k, v := range users {
		v.Username = k
	}
	return users, nil
}

func getUser(username string, passwordHash hashing.ClientHash) (*User, error) {
	users, err := readUsers()
	if err != nil {
		return nil, err
	}
	user, ok := users[username]
	if !ok {
		// Slow down the request, even if the username was invalid
		nilUser.PasswordHash.Verify(passwordHash)
		return nil, errInvalidCredentials
	}
	if !user.PasswordHash.Verify(passwordHash) {
		return nil, errInvalidCredentials
	}
	return user, nil
}

func addUser(username string, password string) error {
	if username == "" || password == "" {
		return fmt.Errorf("username and password must be provided for new user")
	}

	user := User{
		Username:     username,
		PasswordHash: hashing.NewServerHash(password),
	}
	f, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return fmt.Errorf("unable to open users file: %w", err)
	}
	defer f.Close()

	var users = make(map[string]*User)
	if err := json.NewDecoder(f).Decode(&users); err != nil && err != io.EOF {
		return fmt.Errorf("unable to read users file: %w", err)
	}
	users[username] = &user
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("unable to seek file: %w", err)
	}
	if err := json.NewEncoder(f).Encode(&users); err != nil {
		return fmt.Errorf("unable to write users file %w", err)
	}
	return nil
}
