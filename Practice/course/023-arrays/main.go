package main

import "fmt"

/*
Array is a building block in Go for slices
Use slices instead
Array is a numbered sequence of elements of a single type, called the element type
Cannot change the size if an initialized array
 */

func main()  {
	var x [5]int // x is an array with type int inside and size of 5
	fmt.Println(x)
	x[3] = 42 //assign a value to position 3 in array
	fmt.Println(x)
	fmt.Println(len(x)) //length of arrays
}