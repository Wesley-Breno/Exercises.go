package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float64
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	var num2 float64
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	math.Max(num1, num2)
	var num3 float64
	fmt.Print("Enter the third number: ")
	fmt.Scanln(&num3)

	if num1 > num2 && num1 > num3 {
		fmt.Println(num1)
	} else if num2 > num1 && num2 > num3 {
		fmt.Println(num2)
	} else if num3 > num1 && num3 > num2 {
		fmt.Println(num3)
	} else {
		fmt.Println("No higher value specified")
	}
}
