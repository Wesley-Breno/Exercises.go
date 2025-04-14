package main

import "fmt"

func main() {
	numbersMap := make(map[string][]int)

	fmt.Println("Enter a negative value to stop adding numbers")
	for {
		var num int
		fmt.Print("Enter an integer: ")
		fmt.Scan(&num)

		if num < 0 {
			break
		}

		if num >= 0 && num <= 25 {
			numbersMap["0-25"] = append(numbersMap["0-25"], num)
			continue
		}

		if num >= 26 && num <= 50 {
			numbersMap["26-50"] = append(numbersMap["26-50"], num)
			continue
		}

		if num >= 51 && num <= 75 {
			numbersMap["51-75"] = append(numbersMap["51-75"], num)
			continue
		}

		if num >= 76 && num <= 100 {
			numbersMap["76-100"] = append(numbersMap["76-100"], num)
			continue
		}
	}

	fmt.Println("\nResult")
	fmt.Println("Numbers in the range 0-25:", numbersMap["0-25"])
	fmt.Println("Numbers in the range 26-50:", numbersMap["26-50"])
	fmt.Println("Numbers in the range 51-75:", numbersMap["51-75"])
	fmt.Println("Numbers in the range 76-100:", numbersMap["76-100"])
}
