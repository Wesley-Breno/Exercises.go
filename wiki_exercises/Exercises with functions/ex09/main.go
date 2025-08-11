package main

import (
	"fmt"
	"math/rand"
	"time"
)

func craps() bool {
	num1 := rand.Intn(6) + 1
	num2 := rand.Intn(6) + 1
	totalValue := num1 + num2

	if totalValue == 7 || totalValue == 11 {
		return true
	} else if totalValue == 2 || totalValue == 3 || totalValue == 12 {
		return false
	} else {
		point := totalValue
		for {
			num1 = rand.Intn(6) + 1
			num2 = rand.Intn(6) + 1
			totalValue = num1 + num2

			if totalValue == point {
				return true
			} else if totalValue == 7 {
				return false
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if craps() {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost!")
	}
}
