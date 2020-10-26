package main

import "fmt"


//structs allowing to aggregate data of a different type, almost a class and object
// Creating a value of type person below
// composite data structure/complex data structure

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person //emmbedded structure
	first string
	ltk bool
}




func main()  {

	// annonamous struct
	p3 := struct {
		first string
		last string
		age int
	}{
		first: "John",
		last: "Mops",
		age :31,
	}
	fmt.Println(p3)

	p1 := person{
		first: "John",
		last:  "Mops",
		age: 31,
	}
	p2 := person{
		first: "Alex",
		last:  "Bunny",
		age: 31,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p1.age)
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	fmt.Println(p2.age)

	fmt.Println()

	sa1 := secretAgent{
		person: person{ //inner type promoted to outer type
			first: "John",
			last:  "Mops",
			age:   31,
		},
		first: "something",
		ltk:    false,
	}

	fmt.Println(sa1)
	fmt.Println(
		sa1.first,
		sa1.last,
		sa1.age,
		sa1.ltk,
	)
	fmt.Println(sa1.first, sa1.person.first) // if there is a name collision


}