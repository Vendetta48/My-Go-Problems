package main

// Using a composite literal create an array that holds 5 values of type int
// Assign values to each index postion
// Range over the array and print the values out
// Using format printing print out the type of array
import (
	"fmt"
)

func main() {
	a := [5]int{2, 4, 6, 8, 3}

	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", a)
}
