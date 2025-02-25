package main

import (
	"fmt"
)

func Sum(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func main() {
	var ages []int
	for {
		var age int
		fmt.Print("Enter the person's age [0 to stop]: ")
		fmt.Scan(&age)

		if age == 0 {
			break
		}

		ages = append(ages, age)
	}

	average := Sum(ages) / len(ages)

	fmt.Println("Average age:", average)

	fmt.Print("Based on the ages provided, the group is classified as: ")
	if average < 26 {
		fmt.Println("Young")
	} else if average < 61 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Elderly")
	}
}
