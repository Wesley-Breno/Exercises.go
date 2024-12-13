package main

import "fmt"

func main() {
	var countryA float64 = 80000
	var countryB float64 = 200000
	var totalYears float64

	for {
		if countryA > countryB {
			break
		}
		countryA = countryA + (countryA * 3 / 100)
		countryB = countryB + (countryB * 1.5 / 100)
		totalYears += 1
	}

	fmt.Printf("It will take %v years for country A to surpass country B", totalYears)
}
