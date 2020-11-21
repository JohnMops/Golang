/*
Hands on exercise
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results

*/

package main

import "fmt"

func main() {
	x := foo()
	y, z := bar()
	fmt.Println(x, y, z)
}

func foo() int {
	return 10
}

func bar() (int, string) {
	return 10, "hi"
}
