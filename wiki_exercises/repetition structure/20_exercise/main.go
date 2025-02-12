package main

import "fmt"

func main() {
	for {
		var number int
		fmt.Print("Inform the number: ")
		fmt.Scan(&number)

		if number <= 0 || number > 16 {
			fmt.Println("The number entered must be positive and less than 16!")
			continue
		}

		fmt.Printf("%v! =", number)
		multi := 1
		for i := number; i > 0; i-- {
			multi *= i
			if i == number {
				fmt.Printf(" %v", i)
				continue
			}
			fmt.Printf(" x %v", i)
		}
		fmt.Printf(" = %v", multi)
		fmt.Println()
	}
}
