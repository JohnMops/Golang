/*
Create a program that uses a switch statement with no switch expression specified.
 */

package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("this is true")
	case false:
		fmt.Println("this is false")
	}
}