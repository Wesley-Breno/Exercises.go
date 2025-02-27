package main

import "fmt"

func main() {
	var candidateA int
	var candidateB int
	var candidateC int

	var numberOfVoters int
	fmt.Print("Enter the number of voters who will vote: ")
	fmt.Scan(&numberOfVoters)

	for {
		if numberOfVoters == 0 {
			break
		}

		var vote int
		fmt.Println("\n\nEnter the candidate you want to vote for.")
		fmt.Println("[0] - Candidate A")
		fmt.Println("[1] - Candidate B")
		fmt.Println("[2] - Candidate C")
		fmt.Print("Candidate: ")
		fmt.Scan(&vote)

		if vote == 0 {
			candidateA++
			numberOfVoters--
			continue
		}

		if vote == 1 {
			candidateB++
			numberOfVoters--
			continue
		}

		if vote == 2 {
			candidateC++
			numberOfVoters--
			continue
		}

		fmt.Println("Enter the candidate number correctly!")
	}

	fmt.Println("Election results")
	fmt.Println("Candidate A: ", candidateA)
	fmt.Println("Candidate B: ", candidateB)
	fmt.Println("Candidate C: ", candidateC)
}
