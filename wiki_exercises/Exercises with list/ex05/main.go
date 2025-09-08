package main

import "fmt"

func main() {
	vector := []int{}
	even := []int{}
	odd := []int{}

	for i := 0; i < 20; i++ {
		var number int
		fmt.Println()
		fmt.Print("Enter a number: ")
		fmt.Scan(&number)

		vector = append(vector, number)

		if number%2 == 0 {
			even = append(even, number)
		} else {
			odd = append(odd, number)
		}
	}

	fmt.Println("\nTyped numbers:", vector)
	fmt.Println("Even numbers:", even)
	fmt.Println("Odd numbers:", odd)
}
