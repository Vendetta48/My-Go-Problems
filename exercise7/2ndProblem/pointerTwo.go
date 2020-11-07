package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Geralt",
		last:  "Of Rivia",
		age:   50,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.first = "Yennefer"
	p.last = "Of Vengerburg"
	p.age = 100

}

// Create a person struct
// Create a func called changeMe with a person as a parameter
// change the value stored at the person address
// Create a value of type person and print out the value
// in func main call changeMe and print out the value
