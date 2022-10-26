package main

import (
	"fmt"
	"strings"
)

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzæøå")
var alphabetInverse = map[rune]int{}

func init() {
	for k, v := range alphabet {
		alphabetInverse[v] = k
	}
}

func xcrypt(s, key string, smult int) string {
	text := make([]rune, len(s))
	i := 0
	for _, v := range s {
		c := alphabetInverse[v]
		s := alphabetInverse[rune(key[i%len(key)])] * smult
		text[i] = alphabet[(c+s+len(alphabet))%len(alphabet)]
		i++
	}
	return string(text)
}

func normalizeInput(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", ""))
}

func encrypt(s, key string) string {
	sn := normalizeInput(s)
	return strings.ToUpper(xcrypt(sn, key, 1))

}
func decrypt(s, key string) string {
	sn := normalizeInput(s)
	return xcrypt(sn, key, -1)
}

func main() {
	fmt.Println("Oppgave a)")
	fmt.Println(encrypt("Snart helg", "torsk"))

	fmt.Println("Oppgave b)")
	fmt.Println(decrypt("QZQOBVCAFFKSDC", "brus"))
}
