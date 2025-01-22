package main

import "fmt"

func main() {
	var evens int
	var odds int

	for i := 0; i < 10; i++ {
		var num int
		fmt.Printf("\nEnter the value of the %vth number: ", i+1)
		fmt.Scan(&num)

		if num%2 == 0 {
			evens += 1
			continue
		}
		odds += 1
	}

	fmt.Println()
	fmt.Println("Number of evens: ", evens)
	fmt.Println("Number of odds: ", odds)
}
