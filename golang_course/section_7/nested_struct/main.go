package main

import "fmt"

// Structure of creating an item
type item struct {
	produtoID int
	qtde      int
	preco     float64
}

// Struct that will receive the order ID and a slice with all orders.
type pedido struct {
	userID int
	itens  []item
}

// 'Order' method to see the total amount to be paid based on the orders in 'order.items'
func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

// Creating an order and listing the items and their details to show the total to be paid at the end.
func main() {
	pedido := pedido{
		userID: 1, // Order ID
		itens: []item{ // Slice containing all items, quantity and price
			{produtoID: 1, qtde: 2, preco: 12.10},   // order data
			{produtoID: 2, qtde: 1, preco: 23.49},   // order data
			{produtoID: 11, qtde: 100, preco: 3.13}, // order data
		},
	}
	fmt.Printf("Valor total do pedido é %.2f", pedido.valorTotal()) // Calling method to see the order total.
}
