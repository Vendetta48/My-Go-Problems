// Using the code from the previous example, delete a record from your map
package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being Evil`, `Ice cream`, `Sunsets`},
	}
	//fmt.Println(m)

	m[`fleming_Ian`] = []string{`steaks`, `cigars`, `espionage`}

	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
