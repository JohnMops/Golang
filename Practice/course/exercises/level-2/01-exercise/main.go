//Write a program that prints a number in decimal, binary, and hex

package main

import "fmt"



func main()  {
	x := 100
	fmt.Printf("%d\t\t%b\t\t%#x\n", x, x, x)
}