package main

import (
	"fmt"
	"time"
)

func main() {
	var date string
	fmt.Print("Enter a date in the format dd/mm/aaaa: ")
	fmt.Scan(&date)

	_, err := time.Parse("02/01/2006", date)
	if err != nil {
		fmt.Println("Invalid date ;(")
	} else {
		fmt.Println("Valid date :D")
	}
}
