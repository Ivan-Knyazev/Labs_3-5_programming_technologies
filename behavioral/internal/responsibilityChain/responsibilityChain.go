package responsibilitychain

import (
	"fmt"
)

type Order struct {
	Name     string
	Amount   float64
	Quantity int
}

// Handler interface
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(order *Order) bool
}

// Base implementation of the Handler
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *BaseHandler) Handle(order *Order) bool {
	if h.next != nil {
		return h.next.Handle(order)
	}
	return true
}

// Handler for stock check
type StockCheckHandler struct {
	BaseHandler
}

func (h *StockCheckHandler) Handle(order *Order) bool {
	if order.Quantity > 10 {
		fmt.Println("Недостаточно товара на складе.")
		return false
	}
	fmt.Println("Товар в наличии.")
	return h.BaseHandler.Handle(order)
}

// Handler for payment check
type PaymentCheckHandler struct {
	BaseHandler
}

func (h *PaymentCheckHandler) Handle(order *Order) bool {
	if order.Amount < 1000 {
		fmt.Println("Платеж недостаточен.")
		return false
	}
	fmt.Println("Платеж принят.")
	return h.BaseHandler.Handle(order)
}
