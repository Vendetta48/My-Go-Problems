// using iota create 4 constants for the next 4 years. Print the constant values

package main

import (
	"fmt"
)

const (
	a = iota + 2021
	b = iota + 2021
	c = iota + 2021
	d = iota + 2021
)

func main() {
	fmt.Println(a, b, c, d)
}
