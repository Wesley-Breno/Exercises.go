package main

import (
	"fmt"
	"strconv"
	"strings"
)

var MonthsWith31Days = map[int]bool{1: true, 3: true, 5: true, 7: true, 8: true, 10: true, 12: true}
var MonthsWith30Days = map[int]bool{4: true, 6: true, 9: true, 11: true}

var MonthsInFull = []string{
	"january", "february", "march", "april", "may", "june",
	"july", "august", "september", "october", "november", "december",
}

func dateWithMonthInFull(date string) string {
	parts := strings.Split(date, "/")
	if len(parts) != 3 {
		return ""
	}

	// Check if all parts are numeric
	for _, p := range parts {
		if _, err := strconv.Atoi(p); err != nil {
			return ""
		}
	}

	day, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	year, _ := strconv.Atoi(parts[2])

	if month < 1 || month > 12 {
		return ""
	}

	// Validate days per month
	switch {
	case MonthsWith31Days[month] && day >= 1 && day <= 31:
	case MonthsWith30Days[month] && day >= 1 && day <= 30:
	case month == 2:
		if (year%4 == 0 && day >= 1 && day <= 29) || (year%4 != 0 && day >= 1 && day <= 28) {
			// valid
		} else {
			return ""
		}
	default:
		return ""
	}

	return fmt.Sprintf("%d of %s of %d", day, MonthsInFull[month-1], year)
}

func main() {
	fmt.Println(dateWithMonthInFull("1/12/2005"))  // ✅ 1 of december of 2005
	fmt.Println(dateWithMonthInFull("31/02/2023")) // ❌ ""
}
