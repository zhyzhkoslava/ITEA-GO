package main

import (
	"encoding/json"
	"fmt"
)

const refundTransactionType = "refund"

type Order struct {
	ID           string `json:"id"`
	Amount       int    `json:"amount"`
	Transactions []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"transactions"`
}

func json_main() {
	jsonData := `
		[
			{"id": "f9c81316-0bad-4f7c-93df-dd441c5371f2", "amount": 1099, "transactions": [{"id": "43c2f68e-85aa-4e1f-a22c-7e42d27a560a", "type": "auth"}, {"id": "2025c1f3-a97a-4f0d-bc2c-dcbcea63930a", "type": "settle"}]},
			{"id": "8e08894d-0c8b-475c-8686-5ed147cb13f0", "amount": 300, "transactions": [{"id": "7ee3f3e3-de15-4f43-827e-802a5376953f", "type": "auth"}, {"id": "86ae4de9-55d0-4132-b541-fe3e33c6f838", "type": "refund"}]}
		]
	`

	var orders []Order
	err := json.Unmarshal([]byte(jsonData), &orders)
	if err != nil {
		fmt.Println("Помилка розкодування JSON:", err)
		return
	}

	for _, order := range orders {
		for _, transaction := range order.Transactions {
			if transaction.Type == refundTransactionType {

				fmt.Printf("ID ордера з транзакцією 'refund': %s\n", order.ID)
				break
			}
		}
	}
}
