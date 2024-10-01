package main

import "fmt"

func main() {
	var firstNumber int
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&firstNumber)

	var secondNumber int
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&secondNumber)

	if firstNumber > secondNumber {
		fmt.Println("The first number is the largest")
	}

	if secondNumber > firstNumber {
		fmt.Println("The second number is the largest")
	} else {
		fmt.Println("The numbers are equal")
	}
}
