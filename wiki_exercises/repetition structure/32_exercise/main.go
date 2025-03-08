package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter the number whose factorial you want to see: ")
	fmt.Scan(&number)

	factorial := 1

	fmt.Println("Factorial of", number)
	fmt.Printf("%v! = ", number)
	for i := number; i > 0; i-- {
		if i == 1 {
			fmt.Print(i)
		} else {
			fmt.Print(i)
			fmt.Print(" . ")
		}

		factorial *= i
	}

	fmt.Print(" = ", factorial)
}
