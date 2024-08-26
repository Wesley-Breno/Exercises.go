package main

import "fmt"

func main() {
	var number1 int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number1)

	var number2 int
	fmt.Print("Enter other number: ")
	fmt.Scanln(&number2)
	fmt.Println("The sum of the numbers is equal to::", number1+number2)
}
