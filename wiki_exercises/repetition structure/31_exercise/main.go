package main

import "fmt"

func sum(numbers []float64) float64 {
	total := 0.0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	for {
		var values []float64
		count := 1

		for {
			var value float64
			fmt.Printf("\nEnter the value of the %v product [0 - to finish]: $ ", count)
			fmt.Scan(&value)

			if value == 0 {
				break
			}

			values = append(values, value)
			count++
		}

		total := sum(values)
		fmt.Print("\nTotal amount to be paid: $", total)

		var money float64
		fmt.Print("\nEnter the amount of money to pay: ")
		fmt.Scan(&money)

		fmt.Print("\nThe change you will receive is: $", money-total)
		fmt.Println()
		fmt.Println("Starting a new purchase...")
		fmt.Println()
	}
}
