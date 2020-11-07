package main

import "fmt"

// Create a new type vehicle

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
	// Using the vehicle, truck, and sedan structs: create a value of type truck and assign values to the fields
	// Create a value of type sedan and assign values to the fields

	truck1 := truck{
		vehicle: vehicle{

			color: "red",
			doors: 2,
		},
		fourWheel: true,
	}
	fmt.Println(truck1)

	car := sedan{
		vehicle: vehicle{
			color: "blue",
			doors: 4,
		},
		luxury: true,
	}
	fmt.Println(car)

	fmt.Println(truck1.fourWheel, truck1.color, truck1.doors)
	fmt.Println(car.color, car.doors, car.luxury)
}
