package main

import (
	"fmt"
	"strconv"
)

func ConvertHours(hour, minute int) []string {
	period := ""

	if 12 <= hour && hour < 24 {
		if hour != 12 {
			hour = hour - 12
		}
		period = "P"
	} else if 1 <= hour && hour <= 11 || hour == 24 {
		if hour == 24 {
			hour = hour - 12
		}
		period = "A"
	}

	return []string{strconv.Itoa(hour), strconv.Itoa(minute), period}
}

func ShowHours(hour, minute, period string) {
	if len(hour) == 1 {
		hour = fmt.Sprintf("0%v", hour)
	}

	if len(minute) == 1 {
		minute = fmt.Sprintf("0%v", minute)
	}

	fmt.Printf("\n\nConverted time: %v:%v %vM", hour, minute, period)
}

func main() {
	var hour int
	fmt.Print("Enter the hour (01-24): ")
	fmt.Scan(&hour)

	var minute int
	fmt.Print("\nEnter the minutes (00-59): ")
	fmt.Scan(&minute)

	formattedTime := ConvertHours(hour, minute)
	ShowHours(formattedTime[0], formattedTime[1], formattedTime[2])
}
