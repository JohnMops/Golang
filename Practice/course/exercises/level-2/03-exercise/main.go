//Create TYPED and UNTYPED constants. Print the values of the constants.

package main

import "fmt"

const (
	typed string = "foo"
	untyped = 5
	)


func main()  {
	fmt.Println(typed)
	fmt.Println(untyped)
}