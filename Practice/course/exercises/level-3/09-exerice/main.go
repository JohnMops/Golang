/*
Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

 */

package main

import "fmt"

func main()  {
	favSport := "hockey"
	switch favSport {
	case "skiing":
		fmt.Println("Not favorite")
	case "hockey":
		fmt.Println("favorite")
	}
}