/*
Using a COMPOSITE LITERAL:
- create a SLICE of TYPE int
- assign 10 VALUES
- Range over the slice and print the values out.
- Using format printing
print out the TYPE of the slice

 */

package main

import "fmt"

func main() {
	myslice := []int{1,2,3,4,5,6,7,8,9}
	for i, v := range myslice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", myslice)
}