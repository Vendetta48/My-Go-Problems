// create a variable of type string using a raw string literal and print it

package main

import "fmt"

func main() {
	x := ` test
	this is another line
	this is one more line `

	fmt.Printf(x)
}
