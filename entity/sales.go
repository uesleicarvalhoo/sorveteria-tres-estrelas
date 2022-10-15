package entity

import (
	"time"

	"github.com/google/uuid"
)

type SaleItem struct {
	Name      string  `json:"name"`
	UnitPrice float64 `json:"unit_price"`
	Amount    int     `json:"amount"`
}

type Sale struct {
	ID          uuid.UUID   `json:"id"`
	PaymentType PaymentType `json:"payment_type"`
	Items       []SaleItem  `json:"items" validate:"required,min=1"`
	Total       float64     `json:"total"`
	Description string      `json:"description"`
	Date        time.Time   `json:"date"`
}
