package dto

import (
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
)

type GetSalesByPeriodQuery struct {
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}

type RegisterSalePayload struct {
	Description string            `json:"description"`
	PaymentType sales.PaymentType `json:"payment_type"`
	Items       []sales.CartItem  `json:"items"`
}
