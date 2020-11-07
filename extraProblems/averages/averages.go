package main

import "fmt"

func main() {
	a := 165
	b := 170
	c := 200
	d := 210
	e := 300

	sum := (a + b + c + d + e)
	avg := (sum / 5)

	fmt.Println(sum)
	fmt.Println(avg)
}

// Take the weights of 5 individuals and find the average
// average = sum of all weights / 5
