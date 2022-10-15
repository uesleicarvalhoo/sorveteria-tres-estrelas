package entity

import "time"

type SaleItem struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount int     `json:"amount"`
}

type Sale struct {
	ID          ID          `json:"id"`
	PaymentType PaymentType `json:"payment_type"`
	Items       []SaleItem  `json:"items" validate:"required,min=1"`
	Total       float64     `json:"total"`
	Description string      `json:"description"`
	Date        time.Time   `json:"date"`
}
