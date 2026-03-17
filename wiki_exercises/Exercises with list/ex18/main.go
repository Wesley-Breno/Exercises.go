package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculatePercentage(playerVotes int, totalVotes int) float64 {
	if totalVotes == 0 {
		return 0
	}
	return float64(playerVotes) / float64(totalVotes) * 100
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func formatPercentage(value float64) string {
	text := fmt.Sprintf("%.1f", value)
	return text + "%"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	votes := [24]int{}
	totalVotes := 0

	fmt.Println("Survey: Who was the best player?")
	fmt.Println()

	for {
		fmt.Print("Player number (0=end): ")
		input := readLine(reader)

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Enter a value between 1 and 23 or 0 to exit!")
			continue
		}

		if number == 0 {
			break
		}

		if number < 1 || number > 23 {
			fmt.Println("Enter a value between 1 and 23 or 0 to exit!")
			continue
		}

		votes[number]++
		totalVotes++
	}

	resultLines := []string{}
	resultLines = append(resultLines, "Voting result:")
	resultLines = append(resultLines, "")
	resultLines = append(resultLines, fmt.Sprintf("A total of %d votes were counted.", totalVotes))
	resultLines = append(resultLines, "")
	resultLines = append(resultLines, "Player\tVotes\t\t%")

	highestVotes := 0
	bestPlayer := 0
	playersWithHighest := 0

	for player := 1; player <= 23; player++ {
		if votes[player] == 0 {
			continue
		}

		percentage := calculatePercentage(votes[player], totalVotes)
		resultLines = append(
			resultLines,
			fmt.Sprintf("%d\t\t%d\t\t%s", player, votes[player], formatPercentage(percentage)),
		)

		if votes[player] > highestVotes {
			highestVotes = votes[player]
			bestPlayer = player
			playersWithHighest = 1
		} else if votes[player] == highestVotes {
			playersWithHighest++
		}
	}

	if totalVotes == 0 {
		resultLines = append(resultLines, "No valid votes were counted.")
	} else if playersWithHighest > 1 {
		resultLines = append(resultLines, "There was no single player with the highest vote percentage.")
	} else {
		bestPercentage := calculatePercentage(highestVotes, totalVotes)
		resultLines = append(
			resultLines,
			fmt.Sprintf(
				"The best player was number %d, with %d votes, corresponding to %s of the total votes.",
				bestPlayer,
				highestVotes,
				formatPercentage(bestPercentage),
			),
		)
	}

	fmt.Println()
	for _, line := range resultLines {
		fmt.Println(line)
	}

	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println("Error writing result file:", err)
		return
	}
	defer file.Close()

	for _, line := range resultLines {
		_, _ = file.WriteString(line + "\n")
	}
}
