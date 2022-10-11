package apitypes

import "github.com/norgor/security-crypto-tasks/6/2/internal/hashing"

type LoginRequest struct {
	Username     string             `json:"username"`
	PasswordHash hashing.ClientHash `json:"passwordHash"`
}
