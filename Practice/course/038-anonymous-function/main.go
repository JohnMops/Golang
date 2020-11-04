package main

import "fmt"

func main()  {
	func(){
		fmt.Println("Anonymous func run")
	}()
	func(x int){
		fmt.Println("The meaning of life: ", x)
	}(42)
	// func expression

	f := func(){ // func is also a type an can be assigned
		fmt.Println("my first func expression")
	}
	f()
	a := func(x int){ // func is also a type an can be assigned
		fmt.Println("I was born in ", x)
	}
	a(1989)
	fmt.Println()
	// return a func from a func

	s1 := foo()
	fmt.Println(s1)

}

// return a func from a func
func bar() func() int {
	return func() int {
		return 451
	}
}

func foo() string  {
	s := "Hello from foo"
	return s
}