package main

import (
	"fmt"
)

func main() {
	grades := [3]int{97, 85, 93}    //string etc, this is an array with number of elemnts
	grades2 := [...]int{97, 85, 93} // hold as much elements as we tell it
	var students [3]string          // empty array, another way of adding elements
	students[0] = "John"
	students[1] = "Lisa"
	students[2] = "Bob"

	var matrix [3][3]int
	matrix[0] = [3]int{1, 0, 0}
	matrix[1] = [3]int{0, 0, 0}
	matrix[2] = [3]int{1, 1, 0}

	// arrays in Go are vars
	a := [...]int{1, 2, 3}
	b := a // assigning arrays to a var , creates a copy of that array. If you have millions of elements in your array
	// it can slow you program down
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// a way of dealing with that is with "pointers"

	c := [...]int{1, 2, 3}
	d := &c // "&" means , take "d" and point to the same data of "c" without copying it and creating a new array
	d[1] = 5
	fmt.Println(d)
	fmt.Println(c)

	// arrays are fixed sized that have to be known at compile time
	// this is not always convenient so we may use "slices"

	e := []int{1, 2, 3}
	f := e //slices are naturally referenced with pointers without the need of "&"
	f[1] = 6
	fmt.Printf("E Slice: %v\n\n", e)
	fmt.Printf("F Slice: %v\n\n", e)
	fmt.Printf("E Length: %v\n\n", len(e))
	fmt.Printf("E Capacity: %v\n\n", cap(e)) //capacity of the slice in the array

	g := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	h := g[:]
	i := g[3:]
	j := g[:6]
	k := g[3:6]
	g[5] = 55

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades: %v\n", grades2)
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Printf("Array Lenght: %v\n", len(students))
	fmt.Printf("Our Matrix: %v\n", matrix)
	/*
		grade1 := 97
		grade2 := 85
		grade3 := 93

		fmt.Printf("Grade: %v, %v, %v\n", grade1, grade2, grade3)
	*/

}
