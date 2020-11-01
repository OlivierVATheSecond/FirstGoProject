// First go project
// Author: Olivier van Asperen
// Date: 01/11/2020

package main

import (
	"fmt"
	"strings"
)

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func main() {
	fmt.Print("First num: ")
	var a int
	fmt.Scanln(&a)

	fmt.Print("Second num: ")
	var b int
	fmt.Scanln(&b)

	fmt.Print("Opperation(add/subtract):")
	var opp string
	fmt.Scanln(&opp)
	oppLow := strings.ToLower(opp)

	if oppLow == "add" {
		var c int = add(a, b)
		fmt.Println("Result:", c)
	} else if oppLow == "subtract" {
		var c int = subtract(a, b)
		fmt.Println("Result:", c)
	}
}
