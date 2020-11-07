package main

import "fmt"

func main() {
	basic()
	secondFunc(5)
}

func basic() {
	fmt.Println("basic func")
}

func secondFunc(f int) {
	fmt.Println("My number is", f)

}



// Create a basic func
// Create a func that takes an argument
// Create a func that has a return
// Create a func that has multiple returns
