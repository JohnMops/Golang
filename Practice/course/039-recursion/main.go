package main

import (
	"fmt"
)

func main() {
	var x int
	x = factorial(4)
	fmt.Println(x)
	y := factorialLoop(4)
	fmt.Println(y)
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func factorialLoop(x int) int {
	total := 1
	for ; x > 0; x-- {
		total *= x
	}
	return total
}
