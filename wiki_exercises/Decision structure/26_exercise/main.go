package main

import "fmt"

func main() {
	var strawberryKgs float64
	fmt.Print("Enter the quantity of strawberries in Kgs: ")
	fmt.Scan(&strawberryKgs)

	var appleKgs float64
	fmt.Print("Enter the quantity of apples in Kgs: ")
	fmt.Scan(&appleKgs)

	// Calculating the price for strawberries
	var totalStrawberry float64
	if strawberryKgs <= 5 {
		totalStrawberry = strawberryKgs * 2.50
	} else {
		totalStrawberry = strawberryKgs * 2.20
	}

	// Calculating the price for apples
	var totalApple float64
	if appleKgs <= 5 {
		totalApple = appleKgs * 1.8
	} else {
		totalApple = appleKgs * 1.5
	}

	// Checking if there will be a discount
	if (appleKgs+strawberryKgs) > 8 || (totalApple+totalStrawberry) > 25 {
		total := totalApple + totalStrawberry
		total = total - (total * 10 / 100)
		fmt.Printf("\nTotal to pay with a 10 percent discount: $%v", total)
	} else {
		total := totalApple + totalStrawberry
		fmt.Printf("\nTotal to pay: $%v", total)
	}
}
