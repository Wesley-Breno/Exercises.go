package main

import "fmt"

func main() {
	var salaryPerHour float64
	fmt.Print("Enter your hourly salary: ")
	fmt.Scanln(&salaryPerHour)

	var hoursWorkedPerMonth float64
	fmt.Print("\nEnter how many hours you worked in the month: ")
	fmt.Scanln(&hoursWorkedPerMonth)

	var grossSalary float64 = hoursWorkedPerMonth * salaryPerHour
	var netSalary float64 = grossSalary

	var paidToINSS float64 = netSalary * 8 / 100
	var paidToSindicate float64 = netSalary * 5 / 100

	// 11% income tax discount
	netSalary = netSalary - (netSalary * 11 / 100)

	// 8% INSS discount
	netSalary = netSalary - paidToINSS

	// 5% union discount
	netSalary = netSalary - paidToSindicate

	fmt.Printf("Gross salary: R$ %.2f", grossSalary)
	fmt.Printf("\nTotal paid to INSS: R$ %.2f", paidToINSS)
	fmt.Printf("\nTotal paid to the union: R$ %.2f", paidToSindicate)
	fmt.Printf("\nTotal net salary: R$ %.2f", netSalary)
}
