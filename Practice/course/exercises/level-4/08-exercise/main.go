/*

- Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things.
- Store three records in your map. Print out all of the values, along with their index position in the slice.

`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
`no_dr`, `Being evil`, `Ice cream`, `Sunsets`


*/

package main

import "fmt"

func main() {
	m := map[string][]string{ // map with string keys that holds slices of string values
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice crean", "Sunsets"},
	}

	for i, v := range m { // i and v are 3
		fmt.Println("record for: ", i)
			for j, val := range v { // j and val are 3
				fmt.Printf("\t%d  %v\n", j, val)
			}

	}
}

