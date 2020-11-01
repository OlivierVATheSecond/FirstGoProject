// First go project
// Author: Olivier van Asperen
// Date: 01/11/2020

package main

import (
	"fmt"
)

func main() {
	fmt.Print("First num: ")
	var a int
	fmt.Scanln(&a)

	fmt.Print("Second num: ")
	var b int
	fmt.Scanln(&b)

	var c int = a + b

	fmt.Println("Output:", c)
}
