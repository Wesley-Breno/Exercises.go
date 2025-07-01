package main

import (
	"fmt"
	"strconv"
)

func reverseInteger(num int) string {
	numStr := strconv.Itoa(num)
	newNum := ""
	for i := len(numStr) - 1; i != -1; i-- {
		newNum = newNum + string(numStr[i])
	}
	return newNum
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	response := reverseInteger(num)
	fmt.Println("Response:", response)
}
