package dto

import (
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/transaction"
)

type GetTransactionByPeriodQuery struct {
	StartAt time.Time `query:"startAt"`
	EndAt   time.Time `query:"endAt"`
}

type CreateTransactionPayload struct {
	Type        transaction.Type `json:"type"`
	Value       float32          `json:"value"`
	Description string           `json:"description"`
}
