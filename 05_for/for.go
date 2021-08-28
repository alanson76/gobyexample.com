package main

import (
	"fmt"
	"strings"
)

func main() {

	// for is Go's ONLY looping construct.

	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println(strings.Repeat("-", 20))

	// class initial/condition/after for loop
	for j := 7; j <= 10; j++ {
		fmt.Println(j)
	}
	fmt.Println(strings.Repeat("-", 20))

	// infinite loop and break
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println(strings.Repeat("-", 20))

	// continue to the next iteration of the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
