package main

import "fmt"

func main() {
	x := []int{4, 3, 2, 5, 6, 7, 8}
	total := foo(x...)
	fmt.Println(total)

	y := []int{4, 3, 2, 5, 6, 7, 8}
	total2 := bar(y)
	fmt.Println(total2)

}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v

	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum

}

// In func main create a slice of int that stores random integers
// Create func foo which takes in a veriadic parameter of type int and also it returns an int
