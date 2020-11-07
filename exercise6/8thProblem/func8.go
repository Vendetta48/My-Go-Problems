package main

import "fmt"

func main() {
	x := bar()
	fmt.Println(x())

}

func bar() func() int {
	return func() int {
		return 451
	}
}

// Create a func which returns a func
// Assign the returned func to a variable
// call the returned func
