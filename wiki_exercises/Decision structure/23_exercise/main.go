package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float64
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	var num2 float64
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	var operation string
	fmt.Println("\nChoose the letter of the operation you want to perform: ")
	fmt.Println("a. Even or Odd")
	fmt.Println("b. Positive or Negative")
	fmt.Println("c. Integer or Decimal")
	fmt.Print("Choice: ")
	fmt.Scan(&operation)

	if operation == "a" || operation == "b" || operation == "c" {
		if operation == "a" {
			fmt.Println()

			if int(num1)%2 == 0 {
				fmt.Printf("The number %v is even\n", num1)
			} else {
				fmt.Printf("The number %v is odd\n", num1)
			}

			if int(num2)%2 == 0 {
				fmt.Printf("The number %v is even\n", num2)
			} else {
				fmt.Printf("The number %v is odd\n", num2)
			}

		} else if operation == "b" {
			fmt.Println()

			if num1 >= 0 {
				fmt.Printf("The number %v is positive\n", num1)
			} else {
				fmt.Printf("The number %v is negative\n", num1)
			}

			if num2 >= 0 {
				fmt.Printf("The number %v is positive\n", num2)
			} else {
				fmt.Printf("The number %v is negative\n", num2)
			}

		} else if operation == "c" {
			num1Trunc := math.Trunc(num1)
			num2Trunc := math.Trunc(num2)

			if num1 == num1Trunc {
				fmt.Printf("\nThe number %v is integer\n", num1)
			} else {
				fmt.Printf("\nThe number %v is decimal\n", num1)
			}

			if num2 == num2Trunc {
				fmt.Printf("The number %v is integer\n", num2)
			} else {
				fmt.Printf("The number %v is decimal\n", num2)
			}
		}
	} else {
		fmt.Println("The selected option does not exist. Please try again.")
	}
}
