package main

import "fmt"

func sum(values []float64) float64 {
	var total float64
	for _, value := range values {
		total += value
	}
	return total
}

func main() {
	var numberOfCDs int

	fmt.Print("Enter the number of CDs: ")
	fmt.Scan(&numberOfCDs)

	var cdValues []float64

	for {
		if numberOfCDs <= 0 {
			break
		}

		var investedValue float64
		fmt.Printf("Enter the value of the %vª CD: ", numberOfCDs)
		fmt.Scan(&investedValue)

		cdValues = append(cdValues, investedValue)
		numberOfCDs--
	}

	total := sum(cdValues)
	average := total / float64(len(cdValues))
	fmt.Println("\nTotal invested in CDs: ", total)
	fmt.Println("Average amount spent per CD: ", average)
}
