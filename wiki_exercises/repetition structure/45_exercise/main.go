package main

import "fmt"

type Student struct {
	Answers []string
	Score   float64
}

type Question struct {
	QuestionName     string
	OptionsAndValues map[string]string
	CorrectAnswer    string
}

type HighOrLowScore struct {
	StudentID int
	Score     float64
}

func SumFloat64(slice []float64) float64 {
	total := 0.0
	for _, value := range slice {
		total += value
	}
	return total
}

func GetMaxValue(scores []float64) float64 {
	max := 0.0
	for _, score := range scores {
		if score > max {
			max = score
		}
	}
	return max
}

func GetMinValue(scores []float64) float64 {
	min := 0.0
	for i, score := range scores {
		if i == 0 {
			min = score
		}
		if score < min {
			min = score
		}
	}
	return min
}

func main() {
	students := make(map[int]Student)
	studentID := 1
	continueInput := true
	questions := make(map[int]Question)
	options := []string{"A", "B", "C", "D", "E"}

	for i := 1; i <= 10; i++ {
		optionsAndValues := make(map[string]string)

		var questionName string
		fmt.Printf("\nEnter the title for question #%v: ", i)
		fmt.Scan(&questionName)

		for _, option := range options {
			var optionValue string
			fmt.Printf("\nEnter the text for option %v: ", option)
			fmt.Scan(&optionValue)

			optionsAndValues[option] = optionValue
		}

		var correctOption string
		fmt.Print("\nEnter the letter of the correct answer: ")
		fmt.Scan(&correctOption)

		questions[i] = Question{
			QuestionName:     questionName,
			OptionsAndValues: optionsAndValues,
			CorrectAnswer:    correctOption,
		}
	}

	for {
		if !continueInput {
			break
		}

		var answers []string

		for questionID, questionData := range questions {
			fmt.Println()
			fmt.Println("---------------------------------------------------")
			fmt.Printf("\nQuestion #%v\n", questionID)

			fmt.Println(questionData.QuestionName)
			for option, value := range questionData.OptionsAndValues {
				fmt.Printf("\n[ %v ] -> %v", option, value)
			}

			var answer string
			fmt.Print("\nEnter the letter of your answer: ")
			fmt.Scan(&answer)

			answers = append(answers, answer)
		}

		if len(answers) > 0 {
			count := 0
			score := 0.0

			for _, questionData := range questions {
				if questionData.CorrectAnswer == answers[count] {
					score += 1.0
				}
				count += 1
			}

			students[studentID] = Student{Answers: answers, Score: score}
			studentID += 1

			var continueQuestion string
			fmt.Print("\nIs there another student taking the test?")
			fmt.Print("\n[ 1 ] - Yes")
			fmt.Print("\n[ 2 ] - No")
			fmt.Print("\nAnswer: ")
			fmt.Scan(&continueQuestion)

			if continueQuestion == "1" {
				continueInput = true
			} else if continueQuestion == "2" {
				continueInput = false
			} else {
				fmt.Println("Please enter a valid number.")
			}
		}
	}

	allScores := []float64{}
	for _, student := range students {
		allScores = append(allScores, student.Score)
	}

	var highestScore HighOrLowScore
	var lowestScore HighOrLowScore

	for id, student := range students {
		if student.Score == GetMaxValue(allScores) {
			highestScore = HighOrLowScore{
				StudentID: id,
				Score:     student.Score,
			}
		}
		if student.Score == GetMinValue(allScores) {
			lowestScore = HighOrLowScore{
				StudentID: id,
				Score:     student.Score,
			}
		}
	}

	totalStudents := len(students)
	averageScore := SumFloat64(allScores) / float64(totalStudents)

	fmt.Println("Student with highest score: ", highestScore.StudentID)
	fmt.Println("Highest score: ", highestScore.Score)
	fmt.Println("============================")
	fmt.Println("Student with lowest score: ", lowestScore.StudentID)
	fmt.Println("Lowest score: ", lowestScore.Score)
	fmt.Println("============================")
	fmt.Println("Average score: ", averageScore)
	fmt.Println("Total students: ", totalStudents)

	fmt.Println("\nAnswer Key")
	for id, question := range questions {
		fmt.Printf("\nQuestion #%v", id)
		fmt.Println("Correct answer:", question.CorrectAnswer)
	}
}
