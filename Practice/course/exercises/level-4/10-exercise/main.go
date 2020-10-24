/*
Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop

 */


package main

import "fmt"

func main() {
	m := map[string][]string{ // map with string keys that holds slices of string values
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice crean", "Sunsets"},
	}
	m["John"] = []string{"Steak", "Salmon", "Sleeping"}

	delete(m, "John")

	if _, ok := m["John"]; ok {
		delete(m, "no_dr")
	}
	for i, v := range m {
		fmt.Println(i, "likes: ")
			for j, val := range v {
				fmt.Printf("\t%d %v\n" , j, val)
			}
	}
}