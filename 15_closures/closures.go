package main

import (
	"fmt"
	"strings"
)

// Go supports anonymous functions, which can form CLOSURES.
// Anonymous functions are useful when you want to define a function inline without having to name it

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println(strings.Repeat("-", 20))
	newInts := intSeq()
	fmt.Println(newInts())
}
