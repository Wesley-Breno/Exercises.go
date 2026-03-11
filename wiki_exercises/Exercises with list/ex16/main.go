package main

import "fmt"

func main() {
	salesCounts := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	salaryRanges := []string{"$200 - $299", "$300 - $399", "$400 - $499",
		"$500 - $599", "$600 - $699", "$700 - $799",
		"$800 - $899", "$900 - $999", "$1000 and above"}

	for {
		var grossSales float64
		fmt.Print("Enter gross sales amount: $")
		if _, err := fmt.Scan(&grossSales); err != nil {
			fmt.Println("Invalid input. Stopping.")
			break
		}
		if grossSales < 0 {
			break
		}

		salary := 200 + (grossSales * 0.09)
		position := int(salary/100) - 2

		if position >= 0 {
			if position > 8 {
				position = 8
			}
			salesCounts[position]++
		}
	}

	fmt.Println("Number of salespeople in each salary range:")
	for i, count := range salesCounts {
		fmt.Printf("%s: %d\n", salaryRanges[i], count)
	}
}
