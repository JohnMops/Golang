package main

import "fmt"

func main()  {
	xi := []int{1,2,3,4,5,6,7,8,9,10,11,12}
	fmt.Println("All numbers total: ", sum(xi...))
	fmt.Println()

	fmt.Println("Even numbers total: ", even(sum,xi...))
	fmt.Println()

	fmt.Println("Odd numbers total: ", odd(sum, xi...))

}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func odd(sum func(xi ...int) int, xi ...int) int {
	var oi []int
	for _, v := range xi {
		if v % 2 != 0 {
			oi = append(oi, v)
		}
	}
	return sum(oi...)
}

func even(sum func(xi ...int) int, xi ...int) int {
	var ei []int // even numbers slice
	for _, v := range xi {
		if v % 2 ==0 {
			ei = append(ei, v)
		}
	}
	return sum(ei...)
}