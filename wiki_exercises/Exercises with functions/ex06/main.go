package main

import "fmt"

func TotalAmount(installment float64, daysLate int) float64 {
	if daysLate <= 0 {
		return installment
	}

	amountToPay := installment + (installment * 3 / 100)
	interest := 0.0

	for i := daysLate; i != 0; i-- {
		interest += installment * 0.1 / 100
	}

	return amountToPay + interest
}

func main() {
	total := 0.0
	totalOfInstallments := 0.0
	installmentsCount := 0

	for {
		var installmentValue float64
		fmt.Print("\n\nEnter the installment amount: ")
		fmt.Scan(&installmentValue)

		if installmentValue <= 0 {
			break
		}

		var daysLate int64
		fmt.Print("Enter the number of days late: ")
		fmt.Scan(&daysLate)

		total = TotalAmount(installmentValue, int(daysLate))
		fmt.Print("Final amount to be paid: $", total)
		totalOfInstallments += total
		installmentsCount += 1
	}

	fmt.Print("\nDaily Report\n")
	fmt.Print("Total amount paid in installments: $", totalOfInstallments)
	fmt.Print("\nTotal number of installments paid today: ", installmentsCount)
}
