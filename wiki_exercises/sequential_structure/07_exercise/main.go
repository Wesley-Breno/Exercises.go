package main

import (
	"fmt"
	"math"
)

func main() {
	var side float64

	fmt.Print("Enter the side of the square: ")
	fmt.Scanln(&side)

	fmt.Printf("Twice the area of the square is: %v", math.Pow(side, 2)*2)
}
