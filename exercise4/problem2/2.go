// Using a composite literal create a slice of type int and assign 10 values
// Range over the slice and print the values out then using format printing print the type of the slice

package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", x)
}
