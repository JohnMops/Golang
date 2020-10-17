package main

import "fmt"

var ( // cannot assign other types to assigned var
	y = 42
	z = "Awesome string" // can use z string = "". This is an interpreted string
	h = `This is a raw string and will include everything: "String within a string
	with a new line"`
)


// go is oriented about TYPE
// Static language and not a Dynamic
// variable is declared to gold a Value of a certain type

func main() {
	fmt.Println("Value of y: ", y)
	fmt.Printf("Type of y: %T\n", y)

	fmt.Printf("Type of z: %T\n", z)
	// z = 43 cannot be done since the var with the identifier "z" is a string
	fmt.Printf("Type of h: %T\n", h)
	fmt.Println("Value of h: ", h)

}
