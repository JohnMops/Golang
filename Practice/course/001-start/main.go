package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is a start of a great journey")
	foo()
	fmt.Println("something else")

	for i := 0; i < 100; i++ {
		if i%2 == 0 && i != 0 {
			fmt.Println(i)
		}
	}
	fmt.Println(add(5, 6))
}

func add(x int, y int) int {
	return x + y
}

func foo() {
	fmt.Println("Just a foo function")
}

//control flow:
// 1. sequence
// 2. loop: iterative
// 3. conditional
