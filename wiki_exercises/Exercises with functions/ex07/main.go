package main

import (
	"fmt"
	"strconv"
)

// CountDigits returns the number of digits in an integer
func CountDigits(num int) int {
	numStr := strconv.Itoa(num)
	return len(numStr)
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	lenNum := CountDigits(num)
	fmt.Println("Number of digits in the entered number:", lenNum)
}
