package main

import "fmt"

func main() {
	var meters int
	fmt.Print("Enter the area to be painted in square meters: ")
	fmt.Scanln(&meters)

	var quantityPaintCans int

	if meters%54 == 0 {
		quantityPaintCans = int(meters / 54)
	} else {
		quantityPaintCans = int(meters/54) + 1
	}

	var totalPayable int = quantityPaintCans * 80

	fmt.Printf("Cans of necessary paint: %v\n", quantityPaintCans)
	fmt.Printf("Total payable: %v\n", totalPayable)
}
