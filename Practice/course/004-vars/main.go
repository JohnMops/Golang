package main

import "fmt"

var ( //var block, global scope
	k     = 10 //package level
	z int = 11
)

func main() {
	//declare var and assign a value at the same time using short declaration operator
	// will not work while declaring a global var outside of the function
	x := 43
	fmt.Println("x is: ", x)
	//var keyword declare and assign, works for declaring a global scope var
	var k = 5
	fmt.Println("k is: ", k)
	// declare a var
	var i int //0 by default for int, 0.0 for floats, "" for strings, nil for pointers, functions
	//interfaces, slices, channels and maps
	//assign a value to the var
	fmt.Println("i is: ", i)
	i = 42
	fmt.Println("i is: ", i)
	fmt.Println("z is: ", z)
}
