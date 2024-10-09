package main

import (
	"fmt"
)

func main() {
	var turno string

	fmt.Print("What shift do you study? M-morning | E-Evening | N-Nocturne: ")
	fmt.Scanln(&turno)

	if turno == "M" || turno == "m" {
		fmt.Println("Good Morning!")
	} else if turno == "E" || turno == "e" {
		fmt.Println("Good Afternoon!")
	} else if turno == "N" || turno == "n" {
		fmt.Println("Good Night!")
	} else {
		fmt.Println("Invalid value!")
	}
}
