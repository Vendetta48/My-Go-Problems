package main

import "fmt"

func main() {
	f := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}
	f()

}

// Assign a func to a variable and then call that func
