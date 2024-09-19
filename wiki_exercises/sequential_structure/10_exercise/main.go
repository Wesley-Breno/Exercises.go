package main

import "fmt"

func main() {
	var celsius float64
	fmt.Print("Set the Celsius temperature: ")
	fmt.Scanln(&celsius)

	fmt.Printf("The temperature %v ºC converted to ºF is equivalent to: %.2f", celsius, (celsius*9/5)+32)
}
