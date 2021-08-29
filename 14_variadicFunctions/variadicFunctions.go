package main

import "fmt"

// variadic functions can be called with any number of trailing arguments

// take an arbitary number of ints as arguments
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// If already have multiple arguments in a slice,
	// apply them top a variadic function using func(slice...) like this
	sum(nums...)
}
