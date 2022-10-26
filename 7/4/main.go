package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var alphabet = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ")
var alphabetInverse = map[rune]int{}

func init() {
	for k, v := range alphabet {
		alphabetInverse[v] = k
	}
}

func transform(s string, f func(x int) int) string {
	runes := make([]rune, utf8.RuneCountInString(s))
	i := 0
	for _, v := range s {
		runes[i] = alphabet[f(alphabetInverse[v])]
		i++
	}
	return string(runes)
}

func main() {
	fmt.Println("Brute forcing shift cipher")
	ciphertext := "YÆVFBVBVFRÅVBV"
	for i := 1; i < len(alphabet); i++ {
		text := transform(ciphertext, func(x int) int {
			return ((x - i) + len(alphabet)) % len(alphabet)
		})
		text = strings.ToLower(text)
		fmt.Printf("Nøkkel %02d: %s\n", i, text)
	}
}
