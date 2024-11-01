package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Print("Enter the number: ")
	fmt.Scan(&num)

	numTrunc := math.Trunc(num)

	if num == numTrunc {
		fmt.Println("The number is integer")
	} else {
		fmt.Println("The number is decimal")
	}
}
