package main

import "fmt"

func main() {
	var number int
	fmt.Print("Inform the number: ")
	fmt.Scan(&number)

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
}
