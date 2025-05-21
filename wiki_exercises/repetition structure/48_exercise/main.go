package main

import "fmt"

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var number string
	fmt.Print("Enter a number to see it reversed: ")
	fmt.Scan(&number)

	fmt.Println("Entered number: ", number)
	fmt.Println("Reversed number: ", reverse(number))
}
