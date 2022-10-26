package main

import (
	"fmt"
	"unicode/utf8"
)

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzæøå")
var alphabetInverse = map[rune]int{}

func init() {
	for k, v := range alphabet {
		alphabetInverse[v] = k
	}
}

func f(x int) int {
	return ((11*x - 5) + 29) % 29
}

func fInverse(y int) int {
	return (8*y + 11) % 29
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
	fmt.Println("Oppgave a)")
	for i := 0; i < 28; i++ {
		fmt.Printf("%d ", f(i))
	}
	fmt.Println()

	fmt.Println("\nOppgave c)")
	fmt.Println("Sekvens")
	for i := 0; i < 28; i++ {
		fmt.Printf("%d -> %d\n", f(i), i)
	}
	fmt.Println("Invers funksjon")
	for i := 0; i < 28; i++ {
		fmt.Printf("%d\n", fInverse(f(i)))
	}
	fmt.Println()

	fmt.Println("Oppgave d)")
	fmt.Printf("f(\"alice\")=%s\n", transform("alice", f))
	fmt.Println()

	fmt.Println("Oppgave e)")
	fmt.Printf("f⁻¹(\"siøpbe\")=%s\n", transform("siøpbe", fInverse))
}
