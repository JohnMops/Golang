/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
pass a func into a func as an argument
*/

package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)
}

func foo() int {
	return 42
}

func bar(f func() int) int {
	return f()
}
