// create a for loop using this syntax
// for condition {}
// have it print out the years you have been alive
package main

import (
	"fmt"
)

func main() {
	x := 1995
	for x <= 2020 {
		fmt.Println(x)
		x++
	}
}
