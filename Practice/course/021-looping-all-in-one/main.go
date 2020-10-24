package main

import "fmt"

func main()  {
	for i := 0; i < 100; i ++ {
		if i%2 == 0 {
			fmt.Printf("Even: %d\n", i)
		} else if i%3 == 0 {
			fmt.Printf("Odd: %d\n", i)
		}
	}

}