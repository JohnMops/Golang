package main

import (
	"fmt"
)
//Constants are declared like variables, but with the const keyword.
//Constants can be character, string, boolean, or numeric values.
//Constants cannot be declared using the := syntax.

const ( // typed constant and untyped
	a int32 = 42
	b float64 = 42.78
	c string = "Hello"
)



func main()  {
	// a = 58 // cannot change const
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("\na type: %T\n", a )
	fmt.Printf("b type: %T\n", b)
	fmt.Printf("c type: %T\n", c)
}