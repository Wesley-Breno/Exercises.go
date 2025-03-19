package main

import (
	"fmt"
)

func main() {
	var start int
	var end int
	var number int

	fmt.Print("Enter where the multiplication table should start: ")
	fmt.Scan(&start)

	fmt.Print("Enter where the multiplication table should end: ")
	fmt.Scan(&end)

	if end <= start {
		fmt.Println("The end of the multiplication table must be greater than the start!")
	} else {
		fmt.Print("Enter the number you want to see the multiplication table for: ")
		fmt.Scan(&number)

		for i := start; i < end+1; i++ {
			fmt.Printf("%v x %v = %v\n", i, number, i*number)
		}
	}
}
