/*
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
- first name
- last name
- favorite ice cream flavors
- Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

*/

package main

import "fmt"

type Person struct {
	first_name string
	last_name string
	ice_cream []string
}

func main()  {
	p1 := Person{
		first_name: "John",
		last_name: "Mops",
		ice_cream: []string{"Caramel", "Strawberry", "Chocolate"},
	}
	p2 := Person{
		first_name: "Alex",
		last_name: "Bunny",
		ice_cream: []string{"Chocolate", "Strawberry", "Cookie"},
	}

	fmt.Println(p1.first_name, p1.last_name, "Loves: ")
	for i := range p1.ice_cream {
		fmt.Println(p1.ice_cream[i])
	}

	fmt.Println(p2.first_name,p2.last_name, "Loves: ")
	for i := range p2.ice_cream {
		fmt.Println(p2.ice_cream[i])
	}
}