package main

import "fmt"

func main() {
	var height float64
	fmt.Print("Enter your height: ")
	fmt.Scanln(&height)

	fmt.Println("Your ideal weight is:", (72.7*height)-58)
}
