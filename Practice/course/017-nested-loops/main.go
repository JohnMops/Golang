package main

import "fmt"

func main()  {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Outer loop: %d\n", i)
		for j := 0; j < 3; j++{
			fmt.Printf("\tInner loop: %d\n", j)
		}
	}

	x := 1
	for x < 10{
		fmt.Println(x)
		x ++
	}

	for {
		if x > 9{
			break
		}
		x++

	}
	fmt.Println("done")
}