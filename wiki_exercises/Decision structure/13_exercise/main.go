package main

import "fmt"

func main() {
	var weekDayNumber int
	fmt.Print("Enter the number of the day of the week (1-Sunday, 2-Monday, etc.): ")
	fmt.Scanln(&weekDayNumber)

	weekDays := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	if weekDayNumber <= len(weekDays) && weekDayNumber > 0 {
		fmt.Println("The entered week number corresponds to the day:", weekDays[weekDayNumber-1])
	} else {
		fmt.Println("[ERROR]: Please enter a numeric value that corresponds to a day of the week.")
	}
}
