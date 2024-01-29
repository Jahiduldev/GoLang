package main

import (
	"fmt"
	"math"
)

// Shape is an interface for shapes that have an area.
type Shape interface {
	Area() float64
}

// Circle is a type representing a circle.
type Circle struct {
	Radius float64
}

// Area calculates the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Square is a type representing a square.
type Square struct {
	SideLength float64
}

// Area calculates the area of the square.
func (s Square) Area() float64 {
	return s.SideLength * s.SideLength
}

func main() {
	// Create instances of Circle and Square
	circle := Circle{Radius: 5}
	square := Square{SideLength: 4}

	// Use the Area method through the Shape interface
	printArea(circle)
	printArea(square)
}

func printArea(shape Shape) {
	// Use the common interface method to calculate and print the area
	fmt.Printf("Area: %f\n", shape.Area())
}
