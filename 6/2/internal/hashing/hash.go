package hashing

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"encoding/json"
	"log"

	"golang.org/x/crypto/pbkdf2"
)

const iters = 100_000
const keyLen = 32
const saltLen = 32

type Hash []byte

var hashFunc = sha1.New
var clientSalt [saltLen]byte = *(*[saltLen]byte)([]byte("a_super_secret_salt!do_not_use:)"))

func newSalt() [saltLen]byte {
	var salt [saltLen]byte
	if _, err := rand.Read(salt[:]); err != nil {
		log.Panicf("Unable to generate salt: %v", err)
	}
	return salt
}

type ClientHash struct {
	hash Hash
}

func NewClientHash(password string) ClientHash {
	hash := pbkdf2.Key([]byte(password), clientSalt[:], iters, keyLen, hashFunc)
	return ClientHash{hash}
}

func (c ClientHash) ToServerHash() ServerHash {
	salt := newSalt()
	hash := pbkdf2.Key(c.hash, salt[:], iters, keyLen, hashFunc)
	return ServerHash{salt, hash}
}

func (c ClientHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.hash)
}

func (c *ClientHash) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &c.hash)
}

type ServerHash struct {
	salt [saltLen]byte
	hash Hash
}

func NewServerHash(password string) ServerHash {
	return NewClientHash(password).ToServerHash()
}

func (s ServerHash) MarshalJSON() ([]byte, error) {
	sh := make([]byte, 0, saltLen+keyLen)
	sh = append(sh, s.salt[:]...)
	sh = append(sh, s.hash...)
	return json.Marshal(&sh)
}

func (s *ServerHash) UnmarshalJSON(b []byte) error {
	sh := make([]byte, 0, saltLen+keyLen)
	if err := json.Unmarshal(b, &sh); err != nil {
		return err
	}
	copy(s.salt[:], sh[:saltLen])
	s.hash = sh[saltLen:]
	return nil
}

func (s ServerHash) Verify(hash ClientHash) bool {
	h := pbkdf2.Key(hash.hash, s.salt[:], iters, keyLen, hashFunc)
	return bytes.Equal(h, s.hash)
}
