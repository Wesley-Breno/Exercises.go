package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)

	var numberOfDivisions int
	var primes []int

	for i := 1; i < n+1; i++ {
		count := 0
		for y := 1; y < i+1; y++ {
			if i%y == 0 {
				numberOfDivisions++
				count++
			}
		}

		if count == 2 {
			primes = append(primes, i)
		}
	}

	fmt.Printf("Prime numbers between 1 and %v\n", n)
	fmt.Printf("Prime numbers: %v\n", primes)
	fmt.Printf("Total divisions performed: %v\n", numberOfDivisions)
}
