package main

import "fmt"

func main() {
	var sum float32
	numbers := 4

	for i := 1; i <= numbers; i++ {
		var number float32
		fmt.Printf("Enter number %d: ", i)
		fmt.Scanln(&number)
		sum += number
	}

	average := sum / float32(numbers)
	fmt.Printf("The average of the %d semester grades is: %.2f\n", numbers, average)
}
