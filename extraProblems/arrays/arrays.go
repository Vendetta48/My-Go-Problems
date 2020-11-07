package main

import "fmt"

func main() {
	a := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(a)

	s := [...]string{"Geralt", "Ciri", "Yennefer"}

	fmt.Println(s)

	for i, v := range s {
		fmt.Println(i, v)
	}
}

// Create an array with numbers 0 to 10 and print it out
// Create an array of strings with names
// Range over both records and print it
