package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (s person) speak() {
	fmt.Println("My name is", s.first, s.last, " My age is", s.age)
}

func main() {
	p1 := person{
		first: "Geralt",
		last:  "Of Rivia",
		age:   50,
	}
	fmt.Println(p1)
	p1.speak()

}

// Create a user defined struct with the identifier person, fields: first, last, and age
// attach a method to type person with: the identifier speak, the method should have the person say their name and age
// create a value of type person, call the method from the value of type person
