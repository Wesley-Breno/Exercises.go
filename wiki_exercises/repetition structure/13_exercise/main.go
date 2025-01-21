package main

import "fmt"

func main() {
	var base int
	fmt.Print("Enter the base number: ")
	fmt.Scan(&base)

	var expoent int
	fmt.Print("Enter the expoent number: ")
	fmt.Scan(&expoent)

	result := 1
	for i := expoent; i > 0; i-- {
		result *= base
	}

	fmt.Printf("\nResult: %v", result)
}
