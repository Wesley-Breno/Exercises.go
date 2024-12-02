package main

import (
	"fmt"
	"log"
)

func main() {
	var fuelType string
	fmt.Print("Enter the type of fuel G (Gasoline), A (Alcohol). Fuel type: ")
	fmt.Scan(&fuelType)

	var litersSold float64
	fmt.Print("Enter the quantity of liters sold. Quantity of liters: ")
	fmt.Scan(&litersSold)

	var discount string
	var totalToPayWithDiscount float64

	if fuelType == "G" {
		totalToPay := litersSold * 2.5

		if litersSold <= 20 { // 4% Discount
			totalToPayWithDiscount = totalToPay - (totalToPay * 4 / 100)
			discount = "4%"
		} else { // 6% Discount
			totalToPayWithDiscount = totalToPay - (totalToPay * 6 / 100)
			discount = "6%"
		}

	} else if fuelType == "A" {
		totalToPay := litersSold * 1.9

		if litersSold <= 20 { // 3% Discount
			totalToPayWithDiscount = totalToPay - (totalToPay * 3 / 100)
			discount = "3%"
		} else { // 5% Discount
			totalToPayWithDiscount = totalToPay - (totalToPay * 5 / 100)
			discount = "5%"
		}
	} else {
		log.Fatal("Enter the fuel type correctly.\nG - Gasoline\nA - Alcohol")
	}

	fmt.Println("Fuel type: ", fuelType)
	fmt.Println("Total liters sold: ", litersSold)
	fmt.Printf("\nTotal to pay with a discount of %v: $%v", discount, totalToPayWithDiscount)
}
