package main

import (
	"fmt"
)

var (
	y = 42
)

func main(){ // explore "fmt" doc
	fmt.Println("y is: ", y)
	fmt.Printf("Type of y is: %T\n", y)
	fmt.Printf("Type of y is: %b\n", y)
	fmt.Printf("Type of y is: %x\n", y)
	fmt.Printf("Type of y is: %#x\n", y)

	s := fmt.Sprintf("Type of y is: %#x\n", y)
	fmt.Println(s)


}