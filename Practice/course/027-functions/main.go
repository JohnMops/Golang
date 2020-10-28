package main

import "fmt"

func main()  {
	foo() // pass arguments into the function that it can receive
	bar("John") // argument string into func bar
	s1 := woo("Bunny")
	fmt.Println(s1)
	x, y := mouse("Ian","Fleming")
	fmt.Println(x,y)
}

// func (r receiver) identifier(parameters) (return/s) { code }

func foo() { // this one receives nothing
	fmt.Println("hi from foo func")
}
// everything in Go is PASS BY VALUE
func bar(s string) { // receives string
	fmt.Println("hello", s)
}

func woo(st string) string  {
	return fmt.Sprint("Hello from woo ", st)
}

func mouse(fn, ln string)(string, bool)  {
	a := fmt.Sprint(fn, ln, " says Hi")
	b := false
	return a, b
}