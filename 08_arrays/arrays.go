package main

import "fmt"

func main() {

	// An array is a numbered sequence of elements of a specific length

	// basic creation with zero-values
	var a [5]int
	fmt.Println("empty array: a", a)

	// set and get a value
	a[4] = 4
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// lenth of an array
	fmt.Println("len:", len(a))

	// declaration and initialization
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	// multi-dimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
