// Create and use an anon struct

package main

import "fmt"

func main() {
	favGame := struct {
		genre  string
		rating string
		length string
	}{
		genre:  "Open world RPG",
		rating: "mature",
		length: "long",
	}

	fmt.Println(favGame)
}
