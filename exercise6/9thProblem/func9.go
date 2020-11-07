package main

import "fmt"

func main() {
	x := []int{4, 3, 5, 6, 7, 11, 22, 33}

	y := sum(x)

	fmt.Println(y)
}

func sum(xi []int) int {
	var total int
	for _, v := range xi {
		total += v
	}
	return total

}

func odd(f func(xi []int) int, yi []int) int {
	
}

// Pass a func into a func as an argument. Known as a callback
// Create a slice of int in func main
// create another function that will calculate the sum
// Then create another function which passes in a func as an argument to do something
