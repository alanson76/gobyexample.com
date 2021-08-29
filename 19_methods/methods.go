package main

import "fmt"

type rect struct {
	width, length int
}

// methods with pointer receiver (r *rect)
func (r *rect) area() int {
	return r.width * r.length
}

// methods with value receiver (r rect)
// You may want to use a pointer receiver type to avoid copying on method calls
// or to allow the method to mutate the receiving struct.
func (r rect) perimeter() int {
	return 2*r.width + 2*r.length
}

func main() {

	r := rect{width: 10, length: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perimeter: ", rp.perimeter())
}
