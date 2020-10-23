package main

import "fmt"

func main()  {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
	if 2 == 2 {
		fmt.Println("005")
	}
	if 2 != 2 {
		fmt.Println("006")
	}
	if !(2 == 2) {
		fmt.Println("007")
	}
	if !(2 != 2) {
		fmt.Println("008")
	}

	if x := 42; x == 2 { //scope of x is limited to the condition block
		fmt.Println("001")
	}

	x := 43
	if x == 42 {
		fmt.Println("x is 42")
	} else if x == 41 {
		fmt.Println("x is 41")
	} else {
		fmt.Printf("x is: %v\n", x)
	}


}