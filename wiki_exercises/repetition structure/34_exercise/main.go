package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter the number to check if it is prime or not: ")
	fmt.Scan(&number)

	var divisors []int

	for i := 1; i < number+1; i++ {
		if number%i == 0 {
			divisors = append(divisors, i)
		}
	}

	if len(divisors) == 2 && divisors[0] == 1 && divisors[1] == number {
		fmt.Println("The number is prime")
	} else {
		fmt.Println("The number is NOT prime")
	}
}
