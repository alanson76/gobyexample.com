package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 42})
	fmt.Println(person{name: "Bob", age: 42})
	// Omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})
	// An & prefix yeilds a point to the struct
	fmt.Println(&person{name: "Ann", age: 40})
	// It's Idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("John"))

	// access
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 31
	fmt.Println(sp.age)
	fmt.Println(s.age)

}
