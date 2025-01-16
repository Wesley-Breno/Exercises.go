package main

import "fmt"

func main() {
	var num1 int
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	var num2 int
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	fmt.Println()

	for i := num1; i < num2+1; i++ {
		fmt.Println(i)
	}
}
