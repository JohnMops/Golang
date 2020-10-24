/*
Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

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
	for i, v := range m { // i and v are 3
		fmt.Println("record for: ", i)
		for j, val := range v { // j and val are 3
			fmt.Printf("\t%d  %v\n", j, val)
		}

	}
}

