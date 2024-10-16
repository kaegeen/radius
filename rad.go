package main

import (
	"fmt"
	"math"
)

func main() {
	// Variable to store the radius
	var radius float64

	// Ask the user to input the radius
	fmt.Print("Enter the radius of the circle: ")
	fmt.Scanln(&radius)

	// Calculate the area
	area := math.Pi * radius * radius

	// Print the result
	fmt.Printf("The area of the circle with radius %.2f is %.2f\n", radius, area)
}
