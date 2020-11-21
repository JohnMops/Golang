/*
- create a type SQUARE
- create a type CIRCLE
- attach a method to each that calculates AREA and returns it
circle area= π r 2
square area = L * W
- create a type SHAPE that defines an interface as anything that has the AREA method
- create a func INFO which takes type shape and then prints the area
- create a value of type square
- create a value of type circle
- use func info to print the area of square
- use func info to print the area of circle
*/

package main

import (
	"fmt"
	"math"
)

type square struct {
	radius float64
}
type circle struct {
	length float64
}

func (s square) area() float64 {
	return math.Pi * s.radius * s.radius
}

func (c circle) area() float64 {
	return c.length * c.length
}

type shape interface {
	area() float64
}

func info(x shape) {
	fmt.Println(x.area())
}

func main() {
	s1 := square{
		radius: 5.0,
	}
	c1 := circle{
		length: 6.6,
	}
	info(s1)
	info(c1)
}
