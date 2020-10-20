package main

import (
	"fmt"
)

// uint / int
// uint (unsigned) = always start form 0 and goes up  0-255
//int (signed) = from negative to positive -127 - 127
// byte is an alias for uint8 (8 bits 0-255)
// rune is an alias for int32 (character)
//

var i int
var a float64

func main(){
	k := 4 // default int, cannot assign a float to int 4.5
	j := 5.4 // default float64
	i = 5
	a = 6.7
	fmt.Println("k is: ", k)
	fmt.Println("j is: ", j)
	fmt.Printf("k type: %T\n", k)
	fmt.Printf("j type: %T\n", j)

	fmt.Println("i is: ", i)
	fmt.Println("a is: ", a)

}