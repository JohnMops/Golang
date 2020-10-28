package main

import "fmt"

type person struct {
	first string
	last string
}

func (p person) speak() { // we attached the function "speak" to any value of type secretAgent
	// every value of that type can use this func
	fmt.Println(p.first, p.last, "- person speak")
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameter/s) (return/s) {code}

func (s secretAgent) speak() { // we attached the function "speak" to any value of type secretAgent
	// every value of that type can use this func
	fmt.Println(s.first, s.last, s.ltk, "- secretAgent speak")
}

func bar(h human)  { // secreatagent and person types passed into here
	switch h.(type) {
	case person:
		fmt.Println("I was passed to bar and am a person", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed to bar and am an agent ", h.(secretAgent).first)
	}
	fmt.Println("I am human", h)
}


type human interface { // defines behavior and can give values multiple types
	speak() // anyone that has the value speak is also a human
}

func main()  {
	sa1 := secretAgent{ // this becomes a type of secretAgent and human
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

	p1 := person{
		first: "Buffy",
		last:  "Dog",
	}

	sa1.speak()
	sa2.speak()
	p1.speak()
	bar(p1)
	bar(sa1)
	bar(sa2) // bar is able to take both person and secretAgent because it can take anything that has speak in it



}


