package main

import "fmt"

func main() {
	// Meat prices
	lessThan5KgsPrices := []float64{4.9, 5.9, 6.9}
	moreThan5KgsPrices := []float64{5.8, 6.8, 7.8}
	meatTypes := []string{"Filet Mignon", "Top Sirloin", "Rump Steak"}
	paymentTypes := []string{"Cash", "Store Card"}

	// Choose the type of meat
	var meatType int
	fmt.Println("Enter the type of meat:")
	fmt.Println("[1] - Filet Mignon")
	fmt.Println("[2] - Top Sirloin")
	fmt.Println("[3] - Rump Steak")
	fmt.Print("Selected meat: ")
	fmt.Scan(&meatType)

	if meatType < 1 || meatType > 3 {
		fmt.Println("Invalid meat type!")
		return
	}

	// Choose the amount of meat
	var quantity float64
	fmt.Print("Enter the desired amount of meat (in Kgs): ")
	fmt.Scan(&quantity)

	// Choose the type of payment
	var paymentType int
	fmt.Println("Enter the type of payment:")
	fmt.Println("[1] - Cash")
	fmt.Println("[2] - Store Card")
	fmt.Print("Payment type: ")
	fmt.Scan(&paymentType)

	if paymentType < 1 || paymentType > 2 {
		fmt.Println("Invalid payment type!")
		return
	}

	// Calculate the total price
	var pricePerKg float64
	if quantity <= 5 {
		pricePerKg = lessThan5KgsPrices[meatType-1]
	} else {
		pricePerKg = moreThan5KgsPrices[meatType-1]
	}
	total := pricePerKg * quantity

	// Calculate the discount
	var discount float64
	if paymentType == 2 { // Store Card
		discount = total * 0.05
	}

	// Final amount to pay
	finalAmount := total - discount

	// Display the receipt
	fmt.Println("\n--- Receipt ---")
	fmt.Printf("Type of meat: %s\n", meatTypes[meatType-1])
	fmt.Printf("Amount of meat: %.2f Kgs\n", quantity)
	fmt.Printf("Total price: $ %.2f\n", total)
	fmt.Printf("Payment type: %s\n", paymentTypes[paymentType-1])
	fmt.Printf("Discount: $ %.2f\n", discount)
	fmt.Printf("Amount to pay: $ %.2f\n", finalAmount)
}
