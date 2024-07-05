package main

import "fmt"

func main() {
	// Creating an array of 5 spaces from 1 to 5
	numeros := [...]int{1, 2, 3, 4, 5}

	// Cycling through the indices and spaces of an array
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// Going through the spaces of an array and ignoring the index
	for _, num := range numeros {
		fmt.Println(num)
	}
}
