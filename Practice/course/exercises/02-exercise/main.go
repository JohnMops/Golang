/*
Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
- identifier “x” type int
- identifier “y” type string
- identifier “z” type bool
in func main
- print out the values for each identifier
- The compiler assigned values to the variables. What are these values called? - 0  values

 */

package main

import "fmt"

var (
	x int
	y string
	z bool
)

func main(){
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)


}