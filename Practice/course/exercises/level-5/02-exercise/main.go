/*
Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

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
	fmt.Println()
	m := map[string]Person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for _, v := range m {
		fmt.Println(v.first_name) // like p1.lastname.p1.firstname
		fmt.Println(v.last_name)  // like p2.lastname.p2.firstname
		for j, val := range v.ice_cream { // iterate inside p1 and p2 icecream slice
			fmt.Println(j, val)
		}

	}
}