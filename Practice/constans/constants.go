package main

import "fmt"

//const myConst int16 = 27 // global var can be shadowed inside the function

/*
const (

	a = iota // enumarateed const also and scoped to the const block
	b = iota
	c = iota


	a = iota // enumarateed const also
	b
	c
)

const ( //iota resets since it is a new block
	a2 = iota

	)

const (
	_ = iota + 5 // "_" represents a var that does not take space in the memory that we do not care about
	d
	e
	f
)
*/

const (
	a       = 1
	b       = 3
	isAdmin = 1 << iota
	isHQ
	canSeeDbB

	eu
	apac
)

func main() {
	/*
		var roles byte = isAdmin | canSeeDbB | apac
		fmt.Printf("isAdmin: %b\nisHq: %b\nanSeeDB: %b\neu: %b\napac: %b\n", isAdmin, isHQ, canSeeDbB,eu, apac)
		fmt.Printf("%b\n", roles)

		fmt.Printf("apac?: %v\n", apac & roles == apac)
	*/
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", b)

	//const myConst int = 2 // cannot be changed, has to be assigned at init
	//fmt.Printf("%v, %T\n", a, a)
	/*
		fmt.Printf("%v\n", a)
		fmt.Printf("%v\n", b)
		fmt.Printf("%v\n", c)
		fmt.Printf("%v\n", a2)
	*/
	//fmt.Printf("%v\n", d)
}
