package main

import "fmt"

// Variatic function, receives an undefined number of parameters. It's the same thing as python's *args.
// To define a function as a variatic function, use '...'
func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	// Note that I use the function several times and define a random number of parameters.
	fmt.Printf("Media: %.2f\n", media(7.2, 6.3, 8.2, 5.5, 4.4, 8.4, 9.1))
	fmt.Printf("Media: %.2f\n", media(5.5, 4.4, 8.4, 9.1))
	fmt.Printf("Media: %.2f\n", media(9.1))
	fmt.Printf("Media: %.2f\n", media())
}
