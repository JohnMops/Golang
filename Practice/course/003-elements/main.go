package main

import (
	"fmt"
)

func main() {
	var x int // declare
	x = 10    // assign

	i := 2        // short declaration operator + assign
	y := 100 + 10 // statement: made of expressions. "+" is an operator
	fmt.Printf("I is: %T\n", i)
	fmt.Printf("X is: %T\n", x)
	n, err := fmt.Printf("Y is: %T\n", y)
	fmt.Println("Error : ", err, "\nBytes for this line: ", n)

}
