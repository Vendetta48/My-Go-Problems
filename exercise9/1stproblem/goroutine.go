package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	wg.Done()

}

func bar() {

	for i := 0; i <= 50; i++ {
		fmt.Println(i)
	}
	wg.Done()

}

// in addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
//use waitgroups to make sure each goroutine finishes before moving on
