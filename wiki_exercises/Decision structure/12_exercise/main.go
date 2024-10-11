package main

import "fmt"

func main() {
	var hourlyWage, hoursWorkedInMonth float64

	// Input
	fmt.Print("Enter your hourly wage: $ ")
	fmt.Scanln(&hourlyWage)

	fmt.Print("Enter the number of hours you worked this month: ")
	fmt.Scanln(&hoursWorkedInMonth)

	// Calculations
	grossSalary := hourlyWage * hoursWorkedInMonth

	var incomeTaxDiscount float64
	if grossSalary <= 900 {
		incomeTaxDiscount = 0
	} else if grossSalary <= 1500 {
		incomeTaxDiscount = grossSalary * 5 / 100
	} else if grossSalary <= 2500 {
		incomeTaxDiscount = grossSalary * 10 / 100
	} else {
		incomeTaxDiscount = grossSalary * 20 / 100
	}

	inssDiscount := grossSalary * 10 / 100
	fgts := grossSalary * 11 / 100
	totalDiscounts := incomeTaxDiscount + inssDiscount
	netSalary := grossSalary - totalDiscounts

	// Output
	fmt.Printf("\nGross Salary: $ %.2f", grossSalary)
	fmt.Printf("\n(-) Income Tax: $ %.2f", incomeTaxDiscount)
	fmt.Printf("\n(-) INSS: $ %.2f", inssDiscount)
	fmt.Printf("\nFGTS (11%%): $ %.2f", fgts)
	fmt.Printf("\nTotal Discounts: $ %.2f", totalDiscounts)
	fmt.Printf("\nNet Salary: $ %.2f\n", netSalary)
}
