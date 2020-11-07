// Use code from previous example
package main

import "fmt"

// Create your own type person with the underlying type being a struct
type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	// Create two values of type person

	p1 := person{
		first:    "Ryan",
		last:     "Feeley",
		iceCream: []string{"Chocolate", "Cookie Dough", "Cookies and Cream"},
	}

	p2 := person{
		first:    "Erik",
		last:     "Feeley",
		iceCream: []string{"Cookie Dough", "Cookies and Cream", "Vanilla"},
	}
	// Store the values of type person in a map with the key of last name

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)
	// Print out the values ranging over the slice
	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.iceCream {
			fmt.Println(i, val)
		}
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.iceCream {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.iceCream {
		fmt.Println(i, v)
	}
}
