package main

import (
	"fmt"
	"slices"
)

func main() {
	var numbers []int
	for i := 1; i < 6; i++ {
		var number int
		fmt.Printf("Enter the %vº number: ", i)
		fmt.Scan(&number)

		numbers = append(numbers, number)
	}
	fmt.Println("Entered numbers: ", numbers)
	fmt.Println("Largest number: ", slices.Max(numbers))
}
