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

	var num3 float64
	fmt.Print("Enter the third number: ")
	fmt.Scanln(&num3)

	fmt.Println("The product of twice the first and half the second: ", (num1*2)*(num2/2))
	fmt.Println("The sum of three times the first and third: ", (num1*3)+(num3))
	fmt.Println("The third raised to the cube: ", math.Pow(num3, 3))
}
