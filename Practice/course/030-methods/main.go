package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func main()  {
	sa1 := secretAgent{
		person: person{
			first: "John",
			last:  "Mops",
		},
		ltk:    true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Alex",
			last:  "Bunny",
		},
		ltk:    false,
	}

	sa1.speak()
	sa2.speak()


}

// func (r receiver) identifier(parameter/s) (return/s) {code}

func (s secretAgent) speak() { // we attached the function "speak" to any value of type secretAgent
	// every value of that type can use this func
	fmt.Println(s.first, s.last, s.ltk)
}
