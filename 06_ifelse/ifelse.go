package main

import (
	"fmt"
	"strings"
)

func main() {
	// * There is no ternary if in Go

	//if ... else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	fmt.Println(strings.Repeat("-", 20))

	// if ... else if
	// declare/initialize a variable ; and condition
	if num := 9; num < 0 {
		fmt.Println(num, " is negative")
	} else if num == 0 {
		fmt.Println(num, " is 0")
	} else {
		fmt.Println(num, " is positive")
	}
}
