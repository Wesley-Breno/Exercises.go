package main

import (
	"fmt"
)

func getSum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	var numbers []int
	for i := 1; i < 6; i++ {
		var number int
		fmt.Printf("Enter the %vº number: ", i)
		fmt.Scan(&number)
		numbers = append(numbers, number)
	}

	sum := getSum(numbers)
	fmt.Println("The sum of the numbers: ", sum)
	fmt.Println("The average of the numbers: ", sum/len(numbers))
}
