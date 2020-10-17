package main

//yearsWorkedyearsWorked
import (
	"fmt"
	"github.com/johnmops/Golang/Practice/packages"
)

// declaration in the global scope //

//var i int = 42
//var k float32 = 50

// block vars // can help organize vars

var (
	//name        string = "John"
	//job         string = "Devops"
	yearsWorked int = 1
)

var (
	//lastName    string = "John"
	//company     string = "Cloudzone"
	//companyType string = "Startup"
	year int = 2020
)

/*
var h int = 6 //lowercase var is a package level var
var H int = 7 //upper case var is global variable exposed outside
var theURL string = "https://sample.com"
*/
func main() { // block scope vars are for the block itself
	/*
		//var num int
		//num = 5 // assigment
		//num2 := 5 // initialaztion and assigment
		//var num3 int = 5
		//var i float32 = 5 // granual control over the type, if := is used, go will decide on the type on its own , for example it will not choose float32 but 64 instead
		//fmt.Println(num)
		//fmt.Printf("%v, %T", num3, num3) //print value and type of the variable
		fmt.Printf("%v, %T", yearsWorked, yearsWorked) //print value and type of the variable
		fmt.Println("\n" + name)
	*/

	// conversion to a different type
	// Remember when conversting from float to int, data after "." will be lost
	/*
		var i int = 43
		fmt.Printf("%v, %T\n", i, i)

		var k float32
		k = float32(i) // converts i from int to float32 and assigns to j
		fmt.Printf("%v, %T\n", k, k)
	*/

	// converting to string in Go will translate the int to its byte representation. use "strconv"

	/*
		var i int = 43
		fmt.Printf("%v, %T\n", i, i)

		var k string
		k = strconv.Itoa(i) // converts i from int to float32 and assigns to j
		fmt.Printf("%v, %T\n", k, k)
	*/

	fmt.Println(packages.Add(yearsWorked, year))

}
