package main

import (
	"fmt"
	"time"
)

func main() {
	initialYear := 1995
	currentYear := time.Now().Year()
	initialPercentage := 0.75
	initialSalary := 1000.0

	for {
		if initialYear == currentYear {
			break
		}

		initialYear += 1
		initialPercentage = initialPercentage * 2
		initialSalary = 1000 + (1000 * initialPercentage / 1000)
	}

	fmt.Println("Current salary with percentage increases each year")
	fmt.Println("R$", initialSalary)

	fmt.Println("\n===================================================================")

	var inputSalary float64
	fmt.Print("Enter the initial salary to see a different result: ")
	fmt.Scan(&inputSalary)

	initialYear = 1995
	finalSalary := 0.0
	initialPercentage = 0.75
	for {
		if initialYear == currentYear {
			break
		}

		initialYear += 1
		initialPercentage = initialPercentage * 2
		finalSalary = inputSalary + (inputSalary * initialPercentage / 100)
	}

	fmt.Println("\nCurrent salary with percentage increases each year")
	fmt.Println("R$", finalSalary)
}
