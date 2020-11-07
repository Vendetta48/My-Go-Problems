// Create a MAP of type string which is a person's "last_first" name, and a value of type []string which stores all their favorite things
// Store three records in your map
// Print out all of the values, along with the index position in the slice

package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being Evil`, `Ice cream`, `Sunsets`},
	}
	//fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
