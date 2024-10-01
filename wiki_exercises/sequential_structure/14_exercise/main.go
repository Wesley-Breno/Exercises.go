package main

import "fmt"

func main() {
	var weightPisces float64
	fmt.Print("Enter the weight of the fish: ")
	fmt.Scanln(&weightPisces)

	if weightPisces > 50 {
		var weightExceeded float64 = weightPisces - 50
		var valueFine float64 = weightExceeded * 4

		fmt.Printf("\nExceeded amount of weight: %.2f Kg\n", weightExceeded)
		fmt.Printf("Amount of fine to be paid: R$ %.2f\n", valueFine)
	} else {
		fmt.Println("Did not exceed the weight limit.")
	}
}
