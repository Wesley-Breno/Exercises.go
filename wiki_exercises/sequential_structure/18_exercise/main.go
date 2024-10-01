package main

import "fmt"

func main() {
	var fileSize float64
	fmt.Print("Enter the file size in MB: ")
	fmt.Scanln(&fileSize)

	var speedInternet float64
	fmt.Print("\nEnter the internet speed in Mbps: ")
	fmt.Scanln(&speedInternet)

	var time = fileSize / (speedInternet / 8) / 60

	if time >= 1 {
		fmt.Println("File size: ", fileSize)
		fmt.Println("Internet speed: ", speedInternet)
		fmt.Println("Download time in minutes: ", time)
	} else {
		time = fileSize / (speedInternet / 8)
		fmt.Println("File size: ", fileSize)
		fmt.Println("Internet speed: ", speedInternet)
		fmt.Println("Download time in seconds: ", time)
	}
}
