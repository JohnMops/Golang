package main

import (
	"fmt"
	//"strconv"
)

// String is immutable once created
// Sequence of bytes
// code point in utf-8 is a character
// Decimal , Hex, Binary https://www.rapidtables.com/code/text/ascii-table.html
//each code point is a "rune"


func main()  {
	s := "Hello there" // using backticks `` will create a raw string the will include all the spaces and new lines
	s = "Hi, dude" // can reassign but cannot change the existing string
	fmt.Printf("%v, %T\n", s , s)
	bs := []byte(s) // convert to a slice of bytes
	fmt.Println(bs) // print the bytes representing the ASCI representation of letters

	for i := 0; i < len(s); i++{
		fmt.Printf("%#U ", s[i]) //print out the utf code point of the character
	}

	for i, v := range s{
		fmt.Printf("\nat index position %d hex is %#x", i, v)
	}
	fmt.Println("")
}