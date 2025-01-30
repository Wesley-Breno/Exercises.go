package main

import "fmt"

func main() {
	var quantity int
	fmt.Print("Enter the number of values you will input: ")
	fmt.Scan(&quantity)

	smallest := 0
	largest := 0
	sum := 0
	for i := 1; i <= quantity; i++ {
		var num int
		fmt.Printf("Enter the %vᵗʰ number: ", i)
		fmt.Scan(&num)

		if i == 1 {
			smallest = num
			largest = num
		}

		if num > largest {
			largest = num
		}

		if num < smallest {
			smallest = num
		}

		sum += num
	}

	fmt.Println("\n\nSmallest value entered:", smallest)
	fmt.Println("Largest value entered:", largest)
	fmt.Println("Sum of all values:", sum)
}
