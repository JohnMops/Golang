package main

import (
	"fmt"
)

func main()  {
	xi := []int{2,3,4,5,6,7,8,9}


	x := sum(xi...) // func sum expects ints and not a slice but with "..." you are unfolding the slice into
	// individual ints
	fmt.Println(x)

}

// func (r receiver) identifier(parameter/s) (return/s) {code}

func sum(x ...int) int  { // variadic parameter
	fmt.Println(cap(x))
	fmt.Println(len(x))
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}