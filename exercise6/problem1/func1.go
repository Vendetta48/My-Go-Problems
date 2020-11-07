package main

import "fmt"

func main() {
	fmt.Println(foo(5, 2))
	x, y := bar("Geralt", " Of Rivia")
	fmt.Println(x)
	fmt.Println(" Number of monsters killed :", y)

}

func foo(x int, y int) int {
	total := 0
	total = x + y
	return total

}

func bar(fn string, ln string) (string, int) {
	s := fmt.Sprint(fn, "", ln, " is a monster hunter")
	x := 99
	return s, x
}

// Create a func with the identifier foo that returns an int
// Create a func with the identifier bar that returns an int and a string
// Call both funcs
// print out the results
// Syntax of a func: func (r receiver) identifier(parameters) return(s) {...}
