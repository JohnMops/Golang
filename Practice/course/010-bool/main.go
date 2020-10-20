package main

import "fmt"

var x bool //default flse

func main(){
	y := 1 == 2
	fmt.Printf("%T, %v\n", x, x)
	x = true
	fmt.Printf("%T, %v\n", x, x)
	fmt.Println(y)
	k := 10
	i := 5
	fmt.Println(k == i)
}