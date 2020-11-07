package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	truck1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(truck1)
	fmt.Println(sedan1)

	fmt.Println(truck1.fourWheel)
	fmt.Println(sedan1.luxury)
}

// Create a new type vehicle with the underlying type struct with the fields: doors color
// Create two types truck and sedan, underlying type of each is struct
// embed the vehicle type in both
// give the field four wheel to truck which is set to bool, give sedan luxury also set to bool
// Using the structs: create a value of type truck and sedan and assign values to the fields
// Print out each of these values
// Print out a single field from each of these values
