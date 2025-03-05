package main

import "fmt"

func main() {
	var price float64
	fmt.Print("Enter the price of the bread: $ ")
	fmt.Scan(&price)

	fmt.Println("[Price Table]")
	for i := 1; i < 51; i++ {
		fmt.Printf("[%vª] $ %v\n", i, float64(i)*price)
	}
}
