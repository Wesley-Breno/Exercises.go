package main

import "fmt"

func PositiveOrNegative(arg int) string {
	if arg > 0 {
		return "P"
	}
	return "N"
}

func main() {
	result := PositiveOrNegative(-1)
	fmt.Println("Result:", result)

	result = PositiveOrNegative(0)
	fmt.Println("Result:", result)

	result = PositiveOrNegative(1)
	fmt.Println("Result:", result)
}
