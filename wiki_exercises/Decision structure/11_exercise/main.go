package main

import (
	"fmt"
)

func main() {
	var salary, newSalary, raiseAmount float64
	var raisePercentage int

	fmt.Print("Enter your current salary: ")
	fmt.Scanln(&salary)

	if salary < 280 {
		raisePercentage = 20
		raiseAmount = salary * 20 / 100
		newSalary = salary + (salary * 20 / 100)

	} else if salary >= 280 && salary < 700 {
		raisePercentage = 15
		raiseAmount = salary * 15 / 100
		newSalary = salary + (salary * 15 / 100)

	} else if salary >= 700 && salary < 1500 {
		raisePercentage = 10
		raiseAmount = salary * 10 / 100
		newSalary = salary + (salary * 10 / 100)

	} else {
		raisePercentage = 5
		raiseAmount = salary * 5 / 100
		newSalary = salary + (salary * 5 / 100)
	}

	fmt.Printf("\nSalary before the raise: %.2f\n", salary)
	fmt.Printf("Applied raise percentage: %v\n", raisePercentage)
	fmt.Printf("Raise amount: %.2f\n", raiseAmount)
	fmt.Printf("New salary after the raise: %.2f\n", newSalary)
}
