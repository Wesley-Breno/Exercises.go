package main

import (
	"fmt"
	"math"
)

func main() {
	// Getting the value of the radius
	var radius float64
	fmt.Print("Enter the radius value: ")
	fmt.Scanln(&radius)

	// Calculating the area of ​​the circle
	area := math.Pi * math.Pow(radius, 2)

	// Showing the result
	fmt.Printf("The area of the circle is: %.2f", area)
}
