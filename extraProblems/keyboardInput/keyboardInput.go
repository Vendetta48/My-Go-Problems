package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Your name is " + name)

}

// func numb() {
// 	fmt.Print("Enter your number: ")
// 	var yourNumber int
// 	_, err := fmt.Scanf("%d", &yourNumber)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("A number of %d%% is", yourNumber)
// 	if (yourNumber >= 1) && (yourNumber <= 10) {
// 		fmt.Println("Your number is within the range ")
// 	} else {
// 		fmt.Println("Your number is outside of the range ")
// 	}
// }

// Make a problem that lets the user input a name
// Get a number from the console and check if its between 1-10
