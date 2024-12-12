package main

import "fmt"

func main() {
	for {
		var name string
		fmt.Print("Enter your username: ")
		fmt.Scan(&name)

		if len(name) > 3 {
			break
		}
		fmt.Println("Your name must have more than 3 characters.")
	}

	for {
		var age int
		fmt.Print("Enter your age: ")
		fmt.Scan(&age)

		if age >= 0 && age <= 150 {
			break
		}
		fmt.Println("Your age must be between 0 and 150.")
	}

	for {
		var salary float64
		fmt.Print("Enter your salary: ")
		fmt.Scan(&salary)

		if salary > 0 {
			break
		}
		fmt.Println("Your salary must be greater than 0.")
	}

	for {
		var gender string
		fmt.Print("Enter your gender: ")
		fmt.Scan(&gender)

		if gender == "f" || gender == "m" {
			break
		}
		fmt.Println("Please enter the correct gender. f - female, m - male")
	}

	for {
		var maritalStatus string
		fmt.Print("Enter your marital status: ")
		fmt.Scan(&maritalStatus)

		if maritalStatus == "s" || maritalStatus == "c" || maritalStatus == "v" || maritalStatus == "d" {
			break
		}
		fmt.Println("Please enter your marital status correctly.")
	}

	fmt.Println("All set!")
}
