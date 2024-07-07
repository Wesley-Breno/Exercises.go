package main

import "fmt"

// Anonymous function in the 'sum' variable in the global scope
var soma = func(n1, n2 int) int {
	return n1 + n2
}

func main() {
	fmt.Println(soma(2, 3)) // Using anonymous 'sum' function

	// Anonymous function in variable 'sub' in the scope of main()
	sub := func(n1, n2 int) int {
		return n1 - n2
	}

	fmt.Println(sub(2, 3)) // Using anonymous 'sub' function
}
