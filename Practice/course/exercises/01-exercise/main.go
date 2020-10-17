//Hands-on exercise #1
//Write a program that prints a number in decimal, binary, and hex

package main

import "fmt"

var x int

func main()  {
	x = 100
	fmt.Printf("Decimal of x: %d\nBinary of x: %#x\nHex of x: %x\n", x, x, x)
}