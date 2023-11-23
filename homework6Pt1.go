package main

import (
	"fmt"
	"time"
)

type Customer struct {
	Email string
}

type Order struct {
	Customer  Customer
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewOrder(customer Customer) *Order {
	currentTime := time.Now()
	return &Order{
		Customer:  customer,
		Status:    "initiated",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}

func (o *Order) Process() {
	if o.Status == "initiated" {
		o.Status = "processing"
		o.UpdatedAt = time.Now()
		fmt.Println("Order processing...")
	} else {
		fmt.Println("Cannot process order. Invalid status.")
	}
}

func (o *Order) MarkAsSuccess() {
	if o.Status == "processing" {
		o.Status = "success"
		o.UpdatedAt = time.Now()
		fmt.Println("Order marked as success!")
	} else {
		fmt.Println("Cannot mark order as success. Invalid status.")
	}
}

func (o *Order) MarkAsFail() {
	if o.Status == "processing" {
		o.Status = "fail"
		o.UpdatedAt = time.Now()
		fmt.Println("Order marked as fail!")
	} else {
		fmt.Println("Cannot mark order as fail. Invalid status.")
	}
}

func main() {
	customer := Customer{Email: "example@gmail.com"}
	order := NewOrder(customer)

	fmt.Printf("Order Status: %s\n", order.Status)
	fmt.Printf("Created At: %s\n", order.CreatedAt)
	fmt.Printf("Updated At: %s\n", order.UpdatedAt)

	order.Process()
	fmt.Printf("Order Status: %s\n", order.Status)
	fmt.Printf("Updated At: %s\n", order.UpdatedAt)

	order.MarkAsSuccess()
	fmt.Printf("Order Status: %s\n", order.Status)
	fmt.Printf("Updated At: %s\n", order.UpdatedAt)

	order.MarkAsFail()
	fmt.Printf("Order Status: %s\n", order.Status)
	fmt.Printf("Updated At: %s\n", order.UpdatedAt)
}
