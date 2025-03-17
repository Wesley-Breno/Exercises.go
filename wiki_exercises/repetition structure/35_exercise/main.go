package main

import "fmt"

func isPrime(number int) bool {
	var divisors []int

	for i := 1; i < number+1; i++ {
		if number%i == 0 {
			divisors = append(divisors, i)
		}
	}

	if len(divisors) == 2 && divisors[0] == 1 && divisors[1] == number {
		return true
	}

	return false
}

func main() {
	var number int
	fmt.Print("Enter the number to see the primes: ")
	fmt.Scan(&number)

	for i := 1; i < number+1; i++ {
		if isPrime(i) {
			fmt.Printf("%v ", i)
		}
	}
}
