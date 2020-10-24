package main

import "fmt"

//

func main() {
	// x := type{values} // composite literal
	//x := []int{4, 5, 6, 7, 8, 42}
	// x[8] = 10 // cannot change the length as well
	/*
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x[5])
	 */
	/*
	for i := range x {
		fmt.Println(x[i])
	}

	for i, v := range x {
		fmt.Println(i, v)
	}
	 */
	// slicing a slice
	/*
	fmt.Println(x[:])
	fmt.Println(x[2:])
	fmt.Println(x[:4])
	fmt.Println(x[1:3])

	for i := 0; i < 5; i++ { //without using range
		fmt.Println(x[i])
	}
	 */
	// append to a slice
	/*
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014) // append can take an unlimited amount of values [...T]
	fmt.Println(x)
	y := []int{234, 456, 678, 987}
	fmt.Println(y)
	x = append(x, y...) // add all values from "y"
	fmt.Println(x)


	// delete from a slice
	x = append(x[:2], x[4:]...) // takes values up until 3 and appends values from 4 to the end of the slices
	// this results in taking out the values between 2 and 4
	fmt.Println(x)


	b := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("Index of 6 is: ", b[5])
	fmt.Println("Taking out 6 from the slice ...")
	b = append(b[:5], b[7:]...)
	fmt.Println("New slice without 6 is: \n", b)
	 */


	// make is a built in func make([]T, length, capacity)
	// slice sits on top of an array. When you change the slice, the underneath array
	// is replaced with a new array and new values from the slice, this takes some compute power
	// if know your underneath array size, you can use "make" to create a big enough array so it will
	// not recreate it and save on the power
	/*
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println("The length: ", len(x))
	fmt.Println("Underlying capacity: ", cap(x))
	x[0] = 42
	x[9] = 999
	// x[10] = 123 // cant, out of range
	fmt.Println(x)
	fmt.Println("The length: ", len(x))
	fmt.Println("Underlying capacity: ", cap(x))
	x = append(x, 1,3,4,5,6,7,8,9,0,2,3,1,4)
	fmt.Println(x)
	fmt.Println("The length: ", len(x))
	fmt.Println("Underlying capacity: ", cap(x)) // capacity was doubled to hold new slice
	 */

	// multi dimensional slices

	iceCream := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(iceCream)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "hazelnut"}
	fmt.Println(mp)

	xp := [][]string{iceCream,mp} // a slice of a slice of string, multi dimensional
	fmt.Println(xp)
	fmt.Println()
	fmt.Println()
}

//a slice allows you to group together VALUES of the same type
