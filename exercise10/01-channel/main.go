package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)
	// go func() {
	// 	c <- 42
	// }()

	c <- 42

	fmt.Println(<-c)
}

// Get this code working
// using func literal
// using a buffered channel
