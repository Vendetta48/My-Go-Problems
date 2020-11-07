package main

// Write a program that prints a number in decimal, binary, and hex
import (
	"fmt"
)

func main() {
	x := 48
	fmt.Printf("%d \n", x)
	fmt.Printf("%b \n", x)
	fmt.Printf("%#x \n", x)

}
