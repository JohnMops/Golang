package main

import "fmt"

var a int

type dog int //create a type

var b dog

func main() {
	a = 42
	fmt.Printf("Value of a: %v\n", a)
	fmt.Printf("Type of a: %T\n", a)
	b = 43
	fmt.Printf("Value of b: %v\n", b)
	fmt.Printf("Type of b: %T\n", b) // output: Type of b: main.dog
	a = int(b)                       //convert "b" to type "int" knows as "cast" in other languages
	fmt.Printf("Value of a aftering converting b: %v\n", a)
	fmt.Printf("Type of a after converting b: %T\n", a)


}
