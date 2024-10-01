package main

import "fmt"

func main() {
	var meters int
	fmt.Print("Enter the meters: ")
	fmt.Scanln(&meters)

	if meters < 108 {
		var quantitiesOfCansOf25Reais int

		if meters%21 == 0 {
			quantitiesOfCansOf25Reais = meters / 21
		} else {
			quantitiesOfCansOf25Reais = (meters / 21) + 1
		}

		fmt.Printf("\n%v Cans of 3.61\nTotal payable: R$ %.2f", quantitiesOfCansOf25Reais, quantitiesOfCansOf25Reais*25)
	}

	if meters%108 == 0 {
		var quantityOfCansOf80Reais int = meters / 108

		fmt.Printf("\n%v 181 paint cans\nTotal payable: R$ %.2f", quantityOfCansOf80Reais, quantityOfCansOf80Reais*80)
	} else {
		var quantityOfCansOf80Reais = meters / 108
		var quantityOfCansOf25Reais int

		if meters%21 == 0 {
			quantityOfCansOf25Reais = (meters - (108 * quantityOfCansOf80Reais)) / 21
		} else {
			quantityOfCansOf25Reais = ((meters - (108 * quantityOfCansOf80Reais)) / 21) + 1
		}

		fmt.Printf("\n%v Quantity of 181 cans", quantityOfCansOf80Reais)
		fmt.Printf("\n%v Quantity of 25 cans", quantityOfCansOf25Reais)
		fmt.Printf("\nTotal payable: R$ %.2f", (quantityOfCansOf80Reais*80)+(quantityOfCansOf25Reais*25))
	}
}
