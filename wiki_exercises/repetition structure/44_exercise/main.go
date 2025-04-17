package main

import "fmt"

func main() {
	var candidates = map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
	}

	fmt.Println("\n[ Enter a negative number to finish voting ]\n")

	for {
		fmt.Print("[ 1 ] - Jose\n[ 2 ] - Augusto\n[ 3 ] - Cesar\n[ 4 ] - Mario\n[ 5 ] - Null vote\n[ 6 ] - Blank vote\n")
		var voteID int
		fmt.Print("Enter the candidate ID you want to vote for: ")
		fmt.Scan(&voteID)

		if voteID < 0 {
			break
		}

		if voteID > 0 && voteID < 7 {
			candidates[voteID] += 1
		}
	}

	totalVotes := candidates[1] + candidates[2] + candidates[3] + candidates[4] + candidates[5] + candidates[6]

	fmt.Println("\n\t\tElection Results\n\nCandidate Name | Code | Total Votes")
	fmt.Println("Jose              1          ", candidates[1])
	fmt.Println("Augusto           2          ", candidates[2])
	fmt.Println("Cesar             3          ", candidates[3])
	fmt.Println("Mario             4          ", candidates[4])
	fmt.Println("Null votes        5          ", candidates[5])
	fmt.Println("Blank votes       6          ", candidates[6])
	fmt.Printf("Percentage of null votes: %.2f%%\n", (float64(candidates[5])/float64(totalVotes))*100)
	fmt.Printf("Percentage of blank votes: %.2f%%\n", (float64(candidates[6])/float64(totalVotes))*100)
}
