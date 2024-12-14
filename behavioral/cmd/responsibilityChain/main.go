package main

import (
	responsibilitychain "design-patterns/behavioral/internal/responsibilityChain"
	"fmt"
)

func main() {
	// Create order
	order := &responsibilitychain.Order{Name: "TestProduct", Amount: 1100, Quantity: 5}

	// Create handlers for responsibility chain
	stockHandler := &responsibilitychain.StockCheckHandler{}
	paymentHandler := &responsibilitychain.PaymentCheckHandler{}

	// Set next handler for stockHandler
	stockHandler.SetNext(paymentHandler)

	// Processing the order
	if stockHandler.Handle(order) {
		fmt.Println("Готово! Заказ успешно обработан.")
	} else {
		fmt.Println("Отказ! Заказ не может быть обработан.")
	}
}
