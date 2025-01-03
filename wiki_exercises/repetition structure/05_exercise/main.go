package main

import "fmt"

func main() {
	wantsToExit := "0"
	for {
		// Getting population and growth rate for country A
		var countryA float64
		fmt.Print("Enter the population of country A: ")
		fmt.Scan(&countryA)

		var growthRateA float64
		fmt.Print("Enter the growth rate of country A: ")
		fmt.Scan(&growthRateA)

		// Getting population and growth rate for country B
		var countryB float64
		fmt.Print("Enter the population of country B: ")
		fmt.Scan(&countryB)

		var growthRateB float64
		fmt.Print("Enter the growth rate of country B: ")
		fmt.Scan(&growthRateB)

		var totalYears float64

		// Checking the population of each country and performing the calculation
		if countryA < countryB {
			for {
				countryA = countryA + (countryA * growthRateA / 100)
				countryB = countryB + (countryB * growthRateB / 100)
				totalYears++
				if countryA >= countryB {
					break
				}
			}
			fmt.Printf("\n\nCountry A will need %v years to surpass country B in population.", totalYears)
		} else if countryB < countryA {
			for {
				countryA = countryA + (countryA * growthRateA / 100)
				countryB = countryB + (countryB * growthRateB / 100)
				totalYears++
				if countryB >= countryA {
					break
				}
			}
			fmt.Printf("\n\nCountry B will need %v years to surpass country A in population.", totalYears)
		} else {
			fmt.Println("\nThe population of both countries is the same.")
		}

		// Asking the user if they want to exit
		fmt.Print("\nDo you want to exit? [0]==No [1]==Yes: ")
		fmt.Scan(&wantsToExit)
		if wantsToExit == "1" {
			break
		}
	}
}
