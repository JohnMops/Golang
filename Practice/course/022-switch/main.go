package main

import "fmt"

func main()  { // if no default then prints the condition that is met
	/*
	switch {
	case false:
		fmt.Println("This should not print")
	case true:
		fmt.Println("This should print")
		fallthrough //will execute and also print the next one
	case 10 == 10:
		fmt.Println("This should print too")
		fallthrough
	case 2 == 2:
		fmt.Println("This should print")
		fallthrough
	default:
		fmt.Println("Default print")
	}
	 */

	n := "Bond"
	switch "Bond" {
	case "Not Bond", "bond", "Bond": //multiple cases to evaluate
		fmt.Println("This should not print")
	case n:
		fmt.Println("This should print")
	}
}