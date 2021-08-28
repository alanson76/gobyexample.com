package main

import "fmt"

func main() {

	// var declares 1 or more varialbes
	var a = "hello"
	fmt.Println(a)

	// multiple varialbes at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// inferred type
	var d = true
	fmt.Println(d)

	// zero-valued identifier
	var e int
	fmt.Println(e)

	// shorthand for declaration and initialization
	f := "apple"
	fmt.Println(f)

}
