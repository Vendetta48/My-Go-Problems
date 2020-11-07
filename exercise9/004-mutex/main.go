package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	const global = 20
	counter := 0
	wg.Add(global)

	for i := 0; i < global; i++ {
		go func() {
			mu.Lock()
			a := counter
			a++
			counter = a
			mu.Unlock()
			wg.Done()

		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}

// fix the previous race conditon by using mutex
