package main

import "fmt"

func main() {
	var numbers []float64

	for i := 0; i < 10; i++ {
		var num float64
		fmt.Print("Enter a number: ")
		fmt.Scan(&num)

		numbers = append(numbers, num)
	}

	fmt.Println("The numbers you entered in reverse order are:")

	for i := len(numbers); i > 0; i-- {
		if i == 1 {
			fmt.Println(numbers[i-1])
		} else {
			fmt.Printf("%v, ", numbers[i-1])
		}
	}
}
