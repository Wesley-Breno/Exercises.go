package main

import "fmt"

func main() {
	// Taking the entire value of meters
	var meter int
	fmt.Print("Enter the value in meters: ")
	fmt.Scanln(&meter)

	// Showing the entire value of reported meters converted to centimeters
	fmt.Printf("The value in meters of %d in centimeters is equal to: %.2d", meter, meter*100)
}
