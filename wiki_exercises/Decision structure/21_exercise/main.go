package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("PAIR")
	} else {
		fmt.Println("ODD")
	}
}
