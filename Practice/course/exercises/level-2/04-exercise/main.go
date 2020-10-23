/*
Write a program that
- assigns an int to a variable
- prints that int in decimal, binary, and hex
- shifts the bits of that int over 1 position to the left, and assigns that to a variable
- prints that variable in decimal, binary, and hex
 */

package main

import "fmt"

func main()  {
	var x int = 42
	fmt.Printf("x: %d\t\t%b\t\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("y: %d\t\t%b\t\t%#x\n", y, y, y)
}