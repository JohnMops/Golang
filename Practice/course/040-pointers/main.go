package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) //points to the address in memory where this value is stored
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(*b)  //get the value in the address
	fmt.Println(*&a) // all in one, get address and get the value at the address

	*b = 43
	fmt.Println(a) // b is the address of a , changes the value of that address and changes a
	fmt.Println()

	y := 4
	fmt.Println(change(&y))
	fmt.Println(y)
}

func change(z *int) int {
	*z = 5
	return *z
}
