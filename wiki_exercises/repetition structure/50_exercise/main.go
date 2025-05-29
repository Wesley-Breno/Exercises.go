package main

import "fmt"

func main() {
	var numTerms int
	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&numTerms)

	sum := 0.0
	fmt.Print("H = ")

	for i := 1; i <= numTerms; i++ {
		sum += 1.0 / float64(i)
		if i == numTerms {
			fmt.Printf("1/%v = ", i)
		} else {
			fmt.Printf("1/%v + ", i)
		}
	}

	fmt.Print(sum)
}
