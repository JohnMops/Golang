/*
Create a for loop using this syntax
for { }
Have it print out the years you have been alive.

 */
package main

import "fmt"

func main()  {
	bd := 1989
	for {
		if bd <= 2020 {
			fmt.Println(bd)
			bd++
		} else {
			break
		}
	}
}