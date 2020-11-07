// create a for loop using this syntax: for {}
// Have it print out the years I have been alive again

package main

import (
	"fmt"
)

func main() {
	x := 1995
	for {
		if x > 2020 {
			break
		}
		fmt.Println(x)
		x++

	}
}
