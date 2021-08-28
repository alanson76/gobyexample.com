package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	// Swtich statement express conditionals across many branches

	// basic swicth
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("default")
	}
	fmt.Println(strings.Repeat("-", 20))

	// multiple conditions with comma
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's a weekday")
	}
	fmt.Println(strings.Repeat("-", 20))

	// switch without an expression is an alternate way to express if/else logic.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A type switch compares types instead of values (type interface)
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println(t, " is an unkown type ")
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey")
}
