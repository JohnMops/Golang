package main

import "fmt"

func main() {
	// maps are key:value store
	// allow super fast and efficient lookup
	// unordered list

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])
	fmt.Println(m["John"]) // if the values is not stored in map, defaults to the default of type int/str/tc
	fmt.Println()
	v, ok := m["John"] // print the value of the key and wether it is on the map or not
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Println()
	m["James"] = 40
	n := "James"
	if v, ok := m[n]; ok {
		fmt.Println(n, " is in the map", v)
	} else {
		fmt.Println(n, " is not in the map")
	}
	fmt.Println()
	fmt.Println()

	m["Todd"] = 45 // append to map
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println()
	fmt.Println()

	xi := []int{4, 5, 6, 7} // quick print of index and value in slice
	for i, v := range xi {
		fmt.Println(i, v)
	}
	delete(m, "John") // delete something that does not exist
	delete(m, "Todd") // delete from map
	for k, v := range m {
		fmt.Println(k, v)
	}

	if _, ok := m["Miss Moneypenny"]; ok {
		delete(m, "Miss Moneypenny")
	}
	fmt.Println(m)
}
