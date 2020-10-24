/*
Using the code from the previous example, use SLICING to create the following new slices which are then printed:
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]
 */


package main

import "fmt"

func main() {
	myslice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(myslice[:5])
	fmt.Println(myslice[5:])
	fmt.Println(myslice[2:7])
	fmt.Println(myslice[1:6])
}