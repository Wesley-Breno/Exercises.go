package main

import (
	"fmt"
	"sort"
)

func main() {
	var num1, num2, num3 int

	fmt.Print("Enter the value of the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the value of the second number: ")
	fmt.Scanln(&num2)

	fmt.Print("Enter the value of the third number: ")
	fmt.Scanln(&num3)

	numbers := []int{num1, num2, num3}

	// Sorting numbers in descending order
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})

	fmt.Println("The 3 numbers entered in descending form are equal to: ", numbers)
}
