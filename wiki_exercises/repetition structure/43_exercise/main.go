package main

import "fmt"

var menu = make(map[int]float64)
var orders = make(map[float64]float64)

func main() {
	menu[100] = 1.2
	menu[101] = 1.3
	menu[102] = 1.5
	menu[103] = 1.2
	menu[104] = 1.3
	menu[105] = 1.0

	for {
		fmt.Println("Enter a negative number to quit")
		fmt.Println("================================")

		fmt.Println("Products")
		fmt.Println("Product Name     | Code | Price")
		fmt.Println("Hot Dog          100    $1.20")
		fmt.Println("Simple Bauru     101    $1.30")
		fmt.Println("Bauru with Egg   102    $1.50")
		fmt.Println("Hamburger        103    $1.20")
		fmt.Println("Cheeseburger     104    $1.30")
		fmt.Println("Soda             105    $1.00")

		var code int
		fmt.Print("Enter the product code: ")
		fmt.Scan(&code)

		if code < 0 {
			break
		}

		var valid bool
		for productCode := range menu {
			if code == productCode {
				valid = true
			}
		}

		if valid {
			var quantity float64
			fmt.Print("\n\nEnter the quantity: ")
			fmt.Scan(&quantity)

			if quantity < 0 {
				break
			}

			for productCode, price := range menu {
				if code == productCode {
					orders[quantity] = price
					quantity = 0
				}
			}
		} else {
			fmt.Println("\n\nPlease enter a valid product code.")
		}
	}

	var total float64

	fmt.Println("\n\nPurchase Summary")
	fmt.Println("=========================")

	for quantity, price := range orders {
		fmt.Println("Order...")
		fmt.Println("Quantity:", quantity)
		fmt.Println("Price: $", price)
		total = total + (price * quantity)
	}

	fmt.Println("\n\nTotal amount due: $", total)
}
