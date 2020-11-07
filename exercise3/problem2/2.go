// Print every rune code point of the uppercase alphabet three times
// uppercase alphabet per the ascii table goes from 65-90
// if i'm printing the same output three times I should use loops within loops?
package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for y := 0; y < 3; y++ {
			fmt.Printf("\t %#U \n", i)
		}
	}

}
