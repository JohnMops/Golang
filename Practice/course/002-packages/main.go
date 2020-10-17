package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Can take unlimited number of args at any type as per godoc.org/fmt", 42, true)
	//n , _ := fmt.Println("Can take unlimited number of args at any type as per godoc.org/fmt", 42 , true)
	// ignore a return with "_"
	// fmt for example returns int for number of bytes written and error
	fmt.Println(n)
	fmt.Println(err)
}
