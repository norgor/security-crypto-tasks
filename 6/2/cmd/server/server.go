package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/norgor/security-crypto-tasks/6/2/internal/apitypes"
)

const addr = ":8443"

var add = flag.Bool("add", false, "Add a new user. Needs username and password.")
var username = flag.String("username", "", "Username for user.")
var password = flag.String("password", "", "Password for user.")

func login(w http.ResponseWriter, r *http.Request) {
	var loginRequest apitypes.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	log.Printf("Getting user '%s'", loginRequest.Username)
	user, err := getUser(loginRequest.Username, loginRequest.PasswordHash)
	if err != nil {
		if err == errInvalidCredentials {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid username or password."))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		log.Printf("unable to get user: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Welcome, %s! Have fun!", user.Username)))
}

func main() {
	flag.Parse()

	if *add {
		if err := addUser(*username, *password); err != nil {
			log.Fatalf("Unable to add user: %v", err)
		}
		return
	}

	http.HandleFunc("/login", login)

	log.Printf("Server running on %s", addr)
	err := http.ListenAndServeTLS(addr, "server.crt", "server.key", nil)
	log.Panicf("unable to start HTTPS listener: %v", err)
}
