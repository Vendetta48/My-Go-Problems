// write a program that assigns an int to a variable, prints that int in decimal, binary, and hex
// shifts the bit of that int over 1 position to the left, and assigns that to a variable
// print out that new variable in decimal, binary, and hex

package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Printf("%d \t %b \t %#x \n", x, x, x)

	y := x << 1
	fmt.Printf("%d \t %b \t %#x", y, y, y)
}
