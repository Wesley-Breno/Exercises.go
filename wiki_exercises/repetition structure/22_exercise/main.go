package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)

	var divisible []int

	for i := 1; i < n+1; i++ {
		if n%i == 0 {
			divisible = append(divisible, i)
		}
	}

	if len(divisible) == 2 && divisible[0] == 1 && divisible[1] == n {
		fmt.Println("The number is prime")
	} else {
		fmt.Println("The number is NOT prime")
		fmt.Println("Numbers that the given number is divisible by:")
		for _, number := range divisible {
			fmt.Println(number)
		}
	}
}
