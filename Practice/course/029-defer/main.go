package main

import "fmt"

func main()  {
	defer foo() // defer the execution of a function until the containing function exits. In this case until like 8
	bar()
}

func foo()  {
	fmt.Println("foo")
}

func bar()  {
	fmt.Println("bar")
}