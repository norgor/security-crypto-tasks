package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"log"
	"runtime"
	"sync/atomic"

	"golang.org/x/crypto/pbkdf2"
)

var salt = []byte("Saltet til Ola")
var iter = 2048
var hashFunc = sha1.New
var keyLen = 20

var passChars = []byte("!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~")

func passFromNum(iteration uint64) []byte {
	pcl := uint64(len(passChars))
	var chars []byte

	for i := 0; iteration != 0; i++ {
		ci := iteration % pcl
		chars = append(chars, passChars[ci])
		iteration /= pcl
	}

	return chars
}

func bruteforce(start uint64, step uint64, targetHash []byte, done *atomic.Bool, passCh chan string) {
	for i := start; !done.Load(); i += step {
		p := passFromNum(i)
		h := pbkdf2.Key(p, salt, iter, keyLen, hashFunc)
		if bytes.Equal(h, targetHash) {
			passCh <- string(p)
		}

		if i%10_000 == 0 {
			log.Printf("Iteration: %d - %s", i, string(p))
		}
	}
}

func main() {
	hash, err := hex.DecodeString("ab29d7b5c589e18b52261ecba1d3a7e7cbf212c6")
	if err != nil {
		log.Panicf("Unable to decode hex string: %v", err)
	}

	passCh := make(chan string)
	found := atomic.Bool{}

	numCPU := uint64(runtime.NumCPU())
	for i := uint64(0); i < numCPU; i++ {
		go bruteforce(i, numCPU, hash, &found, passCh)
	}
	pass := <-passCh
	log.Printf("Found password: '%s'", pass)
}
