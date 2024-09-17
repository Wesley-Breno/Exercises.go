package main

import "fmt"

func main() {
	var fahrenheit float64
	fmt.Print("Set the Fahrenheit temperature: ")
	fmt.Scanln(&fahrenheit)

	fmt.Printf("The temperature %v ºF converted to ºC is equivalent to: %.2f", fahrenheit, (fahrenheit-32)*5/9)
}
