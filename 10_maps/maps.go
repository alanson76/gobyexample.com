package main

import "fmt"

func main() {

	// create an empty map
	m := make(map[string]int)

	// add key, value
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	// delete
	delete(m, "k2")
	fmt.Println("map:", m)

	// value, exist? present?
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and initilize
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
