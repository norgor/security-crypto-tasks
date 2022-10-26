package main

import (
	"fmt"
	"math"

	"github.com/fatih/color"
)

func printModTable(x int) {
	pad := int(math.Log10(float64((x-1)*(x-1)))) + 1
	ppad := fmt.Sprintf("%%0%dd ", pad)
	for i := 0; i <= pad; i++ {
		fmt.Printf(" ")
	}
	for i := 1; i < x; i++ {
		fmt.Printf(ppad, i)
	}
	fmt.Println()

	for i := 1; i < x; i++ {
		fmt.Printf(ppad, i)
		for j := 1; j < x; j++ {
			var c color.Attribute
			if i*j%x == 1 {
				c = color.FgRed
			} else {
				c = color.FgWhite
			}
			color.New(c).Printf(ppad, i*j)
		}
		fmt.Println()
	}
}

func findMultiplicativeInverse(a, x int) int {
	for i := 1; i < x; i++ {
		if a*i%x == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("Modulo table for %12")
	printModTable(12)
	fmt.Println()

	fmt.Println("Modulo table for %11")
	printModTable(11)
	fmt.Println()

	fmt.Printf("Multiplicative inverse for %d mod %d: %d\n", 11, 29, findMultiplicativeInverse(11, 29))
}
