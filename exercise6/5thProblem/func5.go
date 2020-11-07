package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (c circle) calcArea() float64 {
	return math.Pi * (c.radius * c.radius)

}

func (s square) calcArea() float64 {
	return s.length * s.length
}

type shape interface {
	calcArea() float64
}

func main() {
	c := circle{
		radius: 4.5,
	}

	s := square{
		length: 8,
	}

	info(c)
	info(s)
}

func info(s shape) {
	fmt.Println(s.calcArea())
}

// create a type square struct use float 64
// create a type circle struct use float 64
// Attach a method to each that calc area and returns it
// Create a type shape that defines an interface as anything that has the area method
// create a func info which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of circle and square
