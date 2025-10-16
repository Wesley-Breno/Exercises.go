package main

import "fmt"

func main() {
	sumOfSquares := 0
	for i := 1; i <= 10; i++ {
		var value int
		fmt.Print("Inform a value: ")
		fmt.Scan(&value)
		sumOfSquares += value * value
	}
	fmt.Println("Sum of squares:", sumOfSquares)
}
