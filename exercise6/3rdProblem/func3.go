package main

import "fmt"

func main() {
	defer foo()

	bar()
}

func foo() {
	fmt.Println(10)
}

func bar() {
	fmt.Println(5)
}
