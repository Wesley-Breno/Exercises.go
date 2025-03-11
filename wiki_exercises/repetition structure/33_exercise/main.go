package main

import (
	"fmt"
	"slices"
)

func CalcAverage(slice []float64) float64 {
	total := 0.0
	for _, value := range slice {
		total += value
	}
	return total / float64(len(slice))
}

func main() {
	var temperatures []float64

	for {
		var temperature float64
		fmt.Print("\nEnter the temperature [-999 to exit]: ")
		fmt.Scan(&temperature)

		if temperature == -999 {
			break
		}

		temperatures = append(temperatures, temperature)
	}

	maxTemp := slices.Max(temperatures)
	minTemp := slices.Min(temperatures)
	average := CalcAverage(temperatures)

	fmt.Println("Highest recorded temperature: ", maxTemp)
	fmt.Println("Lowest recorded temperature: ", minTemp)
	fmt.Println("Average temperature: ", average)
}
