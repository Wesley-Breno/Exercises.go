package main

import "fmt"

func main() {
	var totalTerms int
	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&totalTerms)

	denominator := 1.0
	sum := 0.0

	for numerator := 1; numerator <= totalTerms; numerator++ {
		if numerator == totalTerms {
			fmt.Printf("%v/%v\n", numerator, denominator)
		} else {
			fmt.Printf("%v/%v + ", numerator, denominator)
		}
		sum += float64(numerator) / denominator
		denominator += 2
	}

	fmt.Printf("\nTotal sum of %v terms: %v", totalTerms, sum)
}
