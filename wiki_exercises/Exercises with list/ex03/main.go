package main

import "fmt"

func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	var numbers []int

	for i := 0; i < 4; i++ {
		var num int
		fmt.Print("Enter a number: ")
		fmt.Scan(&num)

		numbers = append(numbers, num)
	}

	sum := sumNumbers(numbers...)
	fmt.Printf("The average of the entered numbers is: %.2f\n", float64(sum)/float64(len(numbers)))
}
