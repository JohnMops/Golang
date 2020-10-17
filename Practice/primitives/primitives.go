package main

import (
	"fmt"
)

func main() {
	/*
		var n bool = true

		k := 1 == 1
		j := 1 == 2


		fmt.Printf("%v, %T\n", n, n)

		fmt.Printf("%v, %T\n", k, k)
		fmt.Printf("%v, %T\n", j, j)
	*/

	/*
		## signed integers ##
		int = unspecified size, default
		int8
		int16
		int32
		int64

		## unsigned integers ##
		n := uint16
		byte

		no uint64
	*/

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	/*
		cannot add different integer types e.g:
		a := int
		b := int8
		To make this works, convert one of the vars to the type of the other
		fmt.Println(a + int(b))
	*/

	/*
		## bit shifting ##
		a := 8 // 2^3
		fmt.Println(a << 3) // 2^6
		fmt.Println(a >> 3) // 2^0
	*/

	/*
		## floating point numbers ##
		n := 3.14
		n = 13.7e72
		n = 2.1E14
		fmt.Printf("%v, %T\n", n, n)
		float64 is default
		float32 can be set
	*/

	/*
		## complex number ##
		complex64
		complex128

		var n complex64 = 1 + 2i // 1 is a real part, 2 is an imaginary part

		fmt.Printf("%v, %T\n" , real(n), real(n))
		fmt.Printf("%v, %T\n" , imag(n), image(n))
		// built in functions to extract the part
		var n complex128 = complex(3,5) // takes real and imag part

		## text type ##

		string // any utf8 character.
		// String is just an alias for a byte, to print the value convert back to string]
		fmt.Printf("%v, %T\n", string(s[2]),s[2])
		//we cannot change the string but can concat them
		s := "this is a string"
		s2 := "also a string
		fmt.Printf("%v, %T\n", s + s2, s + s2

		## rune ##

		//represent utf32 char
		//alias for int32

		r := 'a'
		var r rune := 'b'

	*/
}
