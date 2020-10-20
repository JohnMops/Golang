package main

import (
	"fmt"
	)

//IOTA is a pre-decalred identifier of a var

const (
	a = iota << 1 // binary 0 = 0
	b // binary of shifted by 1 = 10 = 2
	c // binary of shifted by 1 = 100 = 4
)

const (
	d = iota + 2
	e
	f
)

const (
	_ = iota // throw 0 away with "_"
	// kb = 1042
	kb = 1 << (iota * 10)
	//mb = 1024 * kb
	mb = 1 << (iota * 10)
	//gb = 1024 * mb
	gb = 1 << (iota * 10)
)

func main()  {
	x := 2
	y := x << 1 // binary 10 shifted to left == 100 == 4
	fmt.Printf("\na type: %T\n", a )
	fmt.Printf("b type: %T\n", b)
	fmt.Printf("c type: %T\n", c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("decimal: %d\tbinary: %b\n", x , x)
	fmt.Printf("decimal: %d\tbinary: %b\n", y , y)
	fmt.Println("")
	fmt.Println("")
	/*
	kb := 1042
	mb := 1024 * kb
	gb := 1024 * mb
	*/
	fmt.Printf("decimal: %d\t\tbinary: %b\n", kb , kb)
	fmt.Printf("decimal: %d\tbinary: %b\n", mb , mb)
	fmt.Printf("decimal: %d\tbinary: %b\n", gb , gb)

}