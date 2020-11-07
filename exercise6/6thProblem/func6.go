package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is an anon func")

	}()

}

// Build and create an anon func
