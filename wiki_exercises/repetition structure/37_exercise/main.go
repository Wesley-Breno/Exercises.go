package main

import (
	"fmt"
)

type clientStruct struct {
	code   int
	height float64
	weight float64
}

func getSkinniest(clients []clientStruct) clientStruct {
	var person clientStruct
	for i, client := range clients {
		if i == 0 {
			person = clientStruct{
				code:   client.code,
				height: client.height,
				weight: client.weight,
			}
			continue
		}

		if client.weight < person.weight {
			person = clientStruct{
				code:   client.code,
				height: client.height,
				weight: client.weight,
			}
		}
	}
	return person
}

func getFattest(clients []clientStruct) clientStruct {
	var person clientStruct
	for i, client := range clients {
		if i == 0 {
			person = clientStruct{
				code:   client.code,
				height: client.height,
				weight: client.weight,
			}
			continue
		}

		if client.weight > person.weight {
			person = clientStruct{
				code:   client.code,
				height: client.height,
				weight: client.weight,
			}
		}
	}
	return person
}

func getShortest(clients []clientStruct) clientStruct {
	var person clientStruct
	for i, client := range clients {
		if i == 0 {
			person = clientStruct{
				code:   client.code,
				height: client.height,
				weight: client.weight,
			}
			continue
		}

		if client.height < person.height {
			person = clientStruct{
				code:   client.code,
				height: client.height,
				weight: client.weight,
			}
		}
	}
	return person
}

func getTotalHeight(clients []clientStruct) float64 {
	var totalHeight float64
	for _, client := range clients {
		totalHeight += client.height
	}
	return totalHeight
}

func getTotalWeight(clients []clientStruct) float64 {
	var totalWeight float64
	for _, client := range clients {
		totalWeight += client.weight
	}
	return totalWeight
}

func main() {
	var clients []clientStruct

	for {
		fmt.Println("\n====================================")
		var code int
		fmt.Print("\nEnter the code: ")
		fmt.Scan(&code)

		if code == 0 {
			break
		}

		var height float64
		fmt.Print("\nEnter the height: ")
		fmt.Scan(&height)

		var weight float64
		fmt.Print("\nEnter the weight: ")
		fmt.Scan(&weight)

		clientVar := clientStruct{
			code:   code,
			height: height,
			weight: weight,
		}

		clients = append(clients, clientVar)
	}

	if len(clients) == 0 {
		return
	}

	skinniest := getSkinniest(clients)
	fmt.Println("Skinniest person: ")
	fmt.Println("Code:", skinniest.code)
	fmt.Println("Height:", skinniest.height)
	fmt.Println("Weight:", skinniest.weight)

	fmt.Println()

	shortest := getShortest(clients)
	fmt.Println("Shortest person: ")
	fmt.Println("Code:", shortest.code)
	fmt.Println("Height:", shortest.height)
	fmt.Println("Weight:", shortest.weight)

	fmt.Println()

	fattest := getFattest(clients)
	fmt.Println("Fattest person: ")
	fmt.Println("Code:", fattest.code)
	fmt.Println("Height:", fattest.height)
	fmt.Println("Weight:", fattest.weight)

	fmt.Println()

	fmt.Println("Average height:", getTotalHeight(clients)/float64(len(clients)))
	fmt.Println("Average weight:", getTotalWeight(clients)/float64(len(clients)))
}
