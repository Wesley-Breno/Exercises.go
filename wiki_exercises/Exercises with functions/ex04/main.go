package main

import "fmt"

func AddTax(taxRate, cost float64) float64 {
	totalCost := cost + (cost * taxRate / 100)
	return totalCost
}

func main() {
	result := AddTax(10, 100)
	fmt.Println("Cost with tax: $", result)
}
