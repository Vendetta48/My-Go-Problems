package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	const global = 20
	counter := 0
	wg.Add(global)

	for i := 0; i < global; i++ {
		go func() {
			a := counter
			runtime.Gosched()
			a++
			counter = a
			wg.Done()

		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}

// using goroutines create an incrementor program
// launch a bunch of goroutines
// each goroutine should read the counter and store it in a new variable
// increment the counter
