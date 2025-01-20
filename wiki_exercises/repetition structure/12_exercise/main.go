package main

import "fmt"

func main() {
	var num1 int
	fmt.Print("Enter the number: ")
	fmt.Scan(&num1)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%v x %v = %v\n", num1, i, num1*i)
	}
}
