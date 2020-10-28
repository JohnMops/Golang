/*
Create and use an anonymous struct.

 */

package main

import "fmt"

func main()  {
	myPerson := struct {
		firstName string
		lastName string
		age int
	}{
		firstName: "John",
		lastName:  "Mops",
		age:       31,
	}

	fmt.Println("My name is ", myPerson.firstName, " last name ", myPerson.lastName, "I am ", myPerson.age, "years old")
}