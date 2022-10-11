package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/norgor/security-crypto-tasks/6/2/internal/apitypes"
	"github.com/norgor/security-crypto-tasks/6/2/internal/hashing"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: ./client <username> <password>\n")
		return
	}
	username := os.Args[1]
	password := os.Args[2]
	passwordHash := hashing.NewClientHash(password)
	loginRequest := apitypes.LoginRequest{
		Username:     username,
		PasswordHash: passwordHash,
	}
	reader, writer := io.Pipe()
	go func() {
		defer writer.Close()
		if err := json.NewEncoder(writer).Encode(&loginRequest); err != nil {
			panic(fmt.Errorf("unable to write request: %v", err))
		}
	}()

	fmt.Println("Sending request...")
	http.DefaultClient.Timeout = 5 * time.Second
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	res, err := http.Post("https://localhost:8443/login", "application/json", reader)
	if err != nil {
		panic(fmt.Errorf("unable to perform request: %v", err))
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(fmt.Errorf("unable to read response: %v", err))
	}
	fmt.Printf("Server responded with: %s\n", string(data))
}
