package dto

import (
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
)

type GetSalesByPeriodQuery struct {
	StartAt time.Time `query:"startAt"`
	EndAt   time.Time `query:"endAt"`
}

type RegisterSalePayload struct {
	Description string            `json:"description"`
	PaymentType sales.PaymentType `json:"payment_type"`
	Items       []sales.CartItem  `json:"items"`
}
