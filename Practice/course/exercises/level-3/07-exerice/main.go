/*
Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

 */

package main

import "fmt"

func main() {
	s := "Bond"
	if s == "Bond" {
		fmt.Println("Bond")
	} else if s == "James" {
		fmt.Println("James")
	} else {
		fmt.Println("Not Bond or James")
	}
}