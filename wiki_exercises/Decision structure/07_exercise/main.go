package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 float64

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	fmt.Print("Enter the third number: ")
	fmt.Scanln(&num3)

	// Taking the highest value
	var bigger float64
	if num1 >= num2 && num1 >= num3 {
		bigger = num1
	} else if num2 >= num1 && num2 >= num3 {
		bigger = num2
	} else if num3 >= num1 && num3 >= num2 {
		bigger = num3
	}

	// Taking the lowest value
	var lowest float64
	if num1 <= num2 && num1 <= num3 {
		lowest = num1
	} else if num2 <= num1 && num2 <= num3 {
		lowest = num2
	} else if num3 <= num1 && num3 <= num2 {
		lowest = num3
	}

	fmt.Println("\nThe lowest value reported was: ", lowest)
	fmt.Println("The highest value reported was: ", bigger)
}
