package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)
	fmt.Println("The number you entered was:", number)
}
