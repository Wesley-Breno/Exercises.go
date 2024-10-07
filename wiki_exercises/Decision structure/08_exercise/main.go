package main

import (
	"fmt"
)

func main() {
	var product1, product2, product3 float64

	fmt.Print("Enter the value of the first product: ")
	fmt.Scanln(&product1)

	fmt.Print("Enter the value of the second product: ")
	fmt.Scanln(&product2)

	fmt.Print("Enter the value of the third product: ")
	fmt.Scanln(&product3)

	// Taking the lowest product value
	if product1 <= product2 && product1 <= product3 {
		fmt.Println("The 1st product is the cheapest, take it.")
	} else if product2 <= product1 && product2 <= product3 {
		fmt.Println("The 2nd product is the cheapest, take it.")
	} else if product3 <= product1 && product3 <= product2 {
		fmt.Println("The 3rd product is the cheapest, take it.")
	}
}
