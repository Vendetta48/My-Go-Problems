package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)

	receive(c)

	fmt.Println("The program is about to exit")

}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i

	}
	close(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println("The value pulled off from the channel : ", v)

	}
}

// Write a program that puts 100 numbers onto a channel
// pulls the the numbers off the channel and prints them out
