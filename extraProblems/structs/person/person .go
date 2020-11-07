package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	p1 := person{
		first:    "Ryan",
		last:     "Feeley2",
		iceCream: []string{"Chocolate", "Cookie dough", "Cookies and Cream"},
	}
	p2 := person{
		first:    "Erik",
		last:     "Feeley",
		iceCream: []string{"Vanilla", "Cookie dough", "Cookies and Cream"},
	}
	// fmt.Println(p1.first)
	// fmt.Println(p1.last)
	// for i, v := range p1.iceCream {
	// 	fmt.Println(i, v)
	// }
	// fmt.Println(p2.first)
	// fmt.Println(p2.last)
	// for j, val := range p2.iceCream {
	// 	fmt.Println(j, val)
	// }

	m := map[string]person{ // key type is string, value type is person
		p1.last: p1,
		p2.last: p2,
	}
	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)

		for i, val := range v.iceCream {
			fmt.Println(i, val)
		}
		fmt.Println("------")
	}
}

// Create your own struct called person so it can store the following data: first name last name ice cream flavor
// Create two values of type person, print out the values, ranging over the elements in the slice which stores the ice cream
// now store the values of type person in a map with the key of last name
