package dto

import (
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

type GetSalesByPeriodQuery struct {
	StartAt time.Time
	EndAt   time.Time
}

type RegisterSalePayload struct {
	Description string             `json:"description"`
	PaymentType entity.PaymentType `json:"payment_type"`
	Items       []entity.CartItem  `json:"items"`
}
