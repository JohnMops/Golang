//Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
//a. 42
//b. “James Bond”
//c. true
//Now print the values stored in those variables using
//a single print statement
//multiple print statements

package main

import "fmt"

func main()  {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x,y,z)
	fmt.Printf("%v\n%v\n%v\n", x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}