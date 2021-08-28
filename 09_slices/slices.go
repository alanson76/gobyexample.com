package main

import "fmt"

func main() {

	// giving a more powerful interface to sequences than arrays

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("app:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied:", s)

	// slice part [startIndex: endIndex] including startIndex: excluding endIndex
	l := s[2:5]
	fmt.Println("slice1", l)

	l = s[:5]
	fmt.Println("slice2", l)

	l = s[2:]
	fmt.Println("slice3", l)

	// delcare and initialize the slice
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	fmt.Println("2d initial: ", twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, 3)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
		fmt.Println("2d: ", twoD)
	}
}
