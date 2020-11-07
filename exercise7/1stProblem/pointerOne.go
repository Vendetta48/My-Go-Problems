package main

import "fmt"

func main() {
	x := 50
	fmt.Println(x)
	y := &x
	fmt.Println(y)

}

// create a value and assign it to a variable
// print the address of that value
