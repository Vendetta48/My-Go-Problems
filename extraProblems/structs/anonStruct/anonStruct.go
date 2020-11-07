package main

import "fmt"

func main() {

	person := struct {
		first string
		last  string
		age   int
	}{

		first: "Ryan",
		last:  "Feeley",
		age:   25,
	}

	fmt.Println(person)
}

// Create and and use an anon struct
