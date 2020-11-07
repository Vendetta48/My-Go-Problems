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
