//Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

package main

import "fmt"

const (
	_ = iota + 2020
	x
	y
	z
	a
)

const (
	c = 2021 + iota
	v = 2021 + iota
	b = 2021 + iota
	n = 2021 + iota
)

func main()  {
	fmt.Println(x, y, z, a)
	fmt.Println(c, v, b, n)
}