package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// To delete from a slice use append along with slicing
	// Delete from the slice above to get this result: [42, 43, 44, 48, 49, 50, 51]
	// Delete 45, 46 and 47
	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	y := x
	fmt.Println(y)
}

// For problem 6 see this: https://play.golang.org/p/dzxZh4lhEpT
