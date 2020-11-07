package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter your number: ")
	var yourNumber int
	_, err := fmt.Scanf("%d", &yourNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("A number of %d%% is", yourNumber)
	if (yourNumber >= 1) && (yourNumber <= 10) {
		fmt.Println("Your number is within the range ")
	} else {
		fmt.Println("Your number is outside of the range ")
	}
}

// take a number from console and see if its within a range 1-10
