package main

import "fmt"

func SumThreeArguments(n1, n2, n3 int) int {
	return n1 + n2 + n3
}

func main() {
	result := SumThreeArguments(1, 1, 1)
	fmt.Println("Result: ", result)
}
