package main

import "fmt"

func CalculateArithmeticMean(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

func main() {
	var amount int
	fmt.Print("Enter the number of numbers that will be used: ")
	fmt.Scan(&amount)

	var nums []float64
	for i := 0; i < amount; i++ {
		var num float64
		fmt.Print("Enter the number: ")
		fmt.Scan(&num)
		nums = append(nums, num)
	}

	result := CalculateArithmeticMean(nums)
	fmt.Println("Result of the arithmetic mean of the numbers reported: ", result)
}
