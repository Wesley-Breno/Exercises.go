package main

import "fmt"

func main() {
	var salaryPerHour float64
	fmt.Print("Enter your hourly wage: ")
	fmt.Scanln(&salaryPerHour)

	var hoursWorkedPerMonth float64
	fmt.Print("Enter the hours you worked in the month: ")
	fmt.Scanln(&hoursWorkedPerMonth)

	fmt.Println("This month you received a salary equivalent to: ", salaryPerHour*hoursWorkedPerMonth)
}
