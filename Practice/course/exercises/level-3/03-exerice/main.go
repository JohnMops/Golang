/*
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.
 */

package main

import "fmt"

func main()  {
	by := 1989
	for by <= 2020 {
		fmt.Println(by)
		by++
	}
}