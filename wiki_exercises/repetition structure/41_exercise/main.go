package main

import (
	"fmt"
	"log"
)

func main() {
	var interest float64
	var installmentValue float64
	var numberOfInstallments float64

	var debt float64
	var installmentOption int

	fmt.Print("Enter the debt amount: $")
	fmt.Scan(&debt)

	fmt.Println("Choose the number of installments")
	fmt.Println("[1] - 1 Installment")
	fmt.Println("[2] - 3 Installments")
	fmt.Println("[3] - 6 Installments")
	fmt.Println("[4] - 9 Installments")
	fmt.Println("[5] - 12 Installments")
	fmt.Print("Option: ")
	fmt.Scan(&installmentOption)

	switch installmentOption {
	case 1:
		interest = 0.0
		numberOfInstallments = 1
		installmentValue = debt
	case 2:
		interest = debt * 10 / 100
		numberOfInstallments = 3
	case 3:
		interest = debt * 15 / 100
		numberOfInstallments = 6
	case 4:
		interest = debt * 20 / 100
		numberOfInstallments = 9
	case 5:
		interest = debt * 25 / 100
		numberOfInstallments = 12
	default:
		log.Fatalln("Please enter a valid option! >:(")
	}

	installmentValue = (debt + interest) / numberOfInstallments

	fmt.Println("Result")
	fmt.Println("Original debt amount:", debt)
	fmt.Println("Interest amount:", interest)
	fmt.Println("Number of installments:", numberOfInstallments)
	fmt.Println("Installment value:", installmentValue)
}
