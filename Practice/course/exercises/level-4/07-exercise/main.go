/*

Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."

- Range over the records, then range over the data in each record.

 */

package main

import (
	"fmt"
)

func main()  {
	s1 := []string{"James", "Bond", "Shaken, not stirred"}
	s2 := []string{"Miss", "Moneypenny", "Helloooooo, James"}

	s3 := [][]string{s1,s2}
	fmt.Println(s3)

	for i, s := range s3 {
		fmt.Println("record: ", i)
			for j, v := range s {
				fmt.Printf("\t index position: %v \t value: %v \n", j , v)
			}
	}
}