package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Athlete struct {
	name    string
	jumps   []float64
	average float64
}

func getAverage(jumps []float64) float64 {
	var sum float64
	for _, jump := range jumps {
		sum += jump
	}
	return sum / float64(len(jumps))
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func readJump(reader *bufio.Reader, label string) float64 {
	for {
		fmt.Printf("%s: ", label)
		valueStr := strings.ReplaceAll(readLine(reader), ",", ".")
		jump, err := strconv.ParseFloat(valueStr, 64)
		if err == nil {
			return jump
		}
		fmt.Println("Invalid value. Enter a number, e.g. 6.5")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	jumpLabels := []string{"First Jump", "Second Jump", "Third Jump", "Fourth Jump", "Fifth Jump"}

	for {
		fmt.Print("Athlete: ")
		name := readLine(reader)
		if name == "" {
			break
		}

		fmt.Println()

		jumps := make([]float64, 0, 5)
		for i := 0; i < 5; i++ {
			jump := readJump(reader, jumpLabels[i])
			fmt.Printf("%.1f m\n", jump)
			fmt.Println()
			jumps = append(jumps, jump)
		}

		athlete := Athlete{name: name, jumps: jumps, average: getAverage(jumps)}

		fmt.Println("Final Result:")
		fmt.Println()
		fmt.Printf("Athlete: %s\n", athlete.name)
		fmt.Printf("Jumps: %.1f - %.1f - %.1f - %.1f - %.1f\n",
			athlete.jumps[0], athlete.jumps[1], athlete.jumps[2], athlete.jumps[3], athlete.jumps[4])
		fmt.Printf("Average of jumps: %.1f m\n", athlete.average)
		fmt.Println()
	}
}
