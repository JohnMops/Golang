/*
Using a COMPOSITE LITERAL:
- create an ARRAY which holds 5 VALUES of TYPE int
- assign VALUES to each index position.
- Range over the array and print the values out.
- Using format printing
print out the TYPE of the array


 */

package main

import "fmt"

func main() {
	myList := [5]int{1, 2, 3, 4, 5}

	for i, v := range myList {
		fmt.Println("At index: ", i, " value is: ", v)
	}
	fmt.Println()
	fmt.Printf("%T\n", myList)

}