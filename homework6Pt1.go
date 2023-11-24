package main

import (
	"errors"
	"fmt"
	"time"
)

type Status int

const (
	Initiated Status = iota
	Processing
	Success
	Fail
)

func (s Status) String() string {
	statusStrings := [...]string{"initiated", "processing", "success", "fail"}
	if s < Initiated || s > Fail {
		return "unknown"
	}
	return statusStrings[s]
}

type Customer struct {
	Email string
}

type Order struct {
	Customer  Customer
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewOrder(customer Customer) *Order {
	currentTime := time.Now()
	return &Order{
		Customer:  customer,
		Status:    Initiated,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}

func (o *Order) Process() error {
	if o.Status == Initiated {
		o.Status = Processing
		o.UpdatedAt = time.Now()
		fmt.Println("Order processing...")
		return nil
	}
	return errors.New("Cannot process order. Invalid status.")
}

func (o *Order) MarkAsSuccess() error {
	if o.Status == Processing {
		o.Status = Success
		o.UpdatedAt = time.Now()
		fmt.Println("Order marked as success!")
		return nil
	}
	return errors.New("Cannot mark order as success. Invalid status.")
}

func (o *Order) MarkAsFail() error {
	if o.Status == Processing {
		o.Status = Fail
		o.UpdatedAt = time.Now()
		fmt.Println("Order marked as fail!")
		return nil
	}
	return errors.New("Cannot mark order as fail. Invalid status.")
}

func main() {
	customer := Customer{Email: "example@gmail.com"}
	order := NewOrder(customer)

	fmt.Printf("Order Status: %s\n", order.Status)
	fmt.Printf("Created At: %s\n", order.CreatedAt)
	fmt.Printf("Updated At: %s\n", order.UpdatedAt)

	err := order.Process()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Order Status: %s\n", order.Status)
		fmt.Printf("Updated At: %s\n", order.UpdatedAt)
	}

	err = order.MarkAsSuccess()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Order Status: %s\n", order.Status)
		fmt.Printf("Updated At: %s\n", order.UpdatedAt)
	}

	err = order.MarkAsFail()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Order Status: %s\n", order.Status)
		fmt.Printf("Updated At: %s\n", order.UpdatedAt)
	}
}
