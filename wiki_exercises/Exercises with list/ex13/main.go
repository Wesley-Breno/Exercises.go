package main

import (
	"fmt"
	"os"
)

func CalculateAverage(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	var sum float64
	for _, num := range numbers {
		sum += num
	}

	return sum / float64(len(numbers))
}

func main() {
	monthsOfYear := []string{
		"January", "February", "March",
		"April", "May", "June", "July",
		"August", "September", "October", "November",
		"December",
	}

	var monthlyAverages []float64
	var annualAverage float64

	for _, month := range monthsOfYear {
		fmt.Println()
		fmt.Println("-------------------------------------")

		var value float64
		fmt.Printf("Enter the average temperature for %s\nTemperature: ", month)
		_, err := fmt.Fscan(os.Stdin, &value)
		if err != nil {
			fmt.Println("Error reading value:", err)
			return
		}

		monthlyAverages = append(monthlyAverages, value)
	}

	annualAverage = CalculateAverage(monthlyAverages)

	fmt.Println()
	fmt.Println("-------------------------------------")
	fmt.Printf("Months with temperature above the annual average of %.2f:\n", annualAverage)

	for i, temp := range monthlyAverages {
		if temp > annualAverage {
			fmt.Printf("%s: %.2f\n", monthsOfYear[i], temp)
		}
	}
}
