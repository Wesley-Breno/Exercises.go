package main

import (
	"fmt"
	"math"
)

func main() {
	var aValue float64
	fmt.Print("Enter the value of A: ")
	fmt.Scanln(&aValue)

	if aValue != 0 {
		var bValue float64
		fmt.Print("Enter the value of B: ")
		fmt.Scanln(&bValue)

		var cValue float64
		fmt.Print("Enter the value of C: ")
		fmt.Scanln(&cValue)

		delta := math.Pow(bValue, 2) - 4*aValue*cValue

		if delta < 0 {
			fmt.Println("\nThe equation does not have real roots.")
			return
		} else if delta == 0 {
			x := -bValue / (2 * aValue)
			fmt.Println("\nThe equation has only one real root.")
			fmt.Println("Root: ", x)
		} else {
			x1 := (-bValue + math.Sqrt(delta)) / (2 * aValue)
			x2 := (-bValue - math.Sqrt(delta)) / (2 * aValue)
			fmt.Println("\nThe equation has two real roots.")
			fmt.Println("Value x1:", x1)
			fmt.Println("Value x2:", x2)
		}

		fmt.Println("\nValue of delta:", delta)

	} else {
		fmt.Println("Based on the value provided, this will not be a quadratic equation.")
	}
}
