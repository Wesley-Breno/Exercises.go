package main

import "fmt"

func main() {
	for {
		var username string
		fmt.Print("Enter your username: ")
		fmt.Scan(&username)

		var password string
		fmt.Print("Enter your password: ")
		fmt.Scan(&password)

		if username == password {
			fmt.Println("Your username cannot be the same as your password!")
		} else {
			fmt.Println("Ok")
			break
		}
	}
}
