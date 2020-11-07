package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("My name is: ", p.first, p.last, "my age: ", p.age)

}

func saySomething(h human) {
	h.speak()

}
func main() {

	p1 := person{
		first: "Geralt",
		last:  "Of Rivia",
		age:   80,
	}
	saySomething(&p1) // this program will not run when p1 does not have & attached to it
}

// see google doc for instructons
