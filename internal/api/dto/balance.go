package dto

import (
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances"
)

type RegisterBalancePayload struct {
	Value       float32                `json:"value"`
	Description string                 `json:"description"`
	Operation   balances.OperationType `json:"operation"`
}

type GetCashFlowByPeriodQuery struct {
	StartAt time.Time `query:"startAt"`
	EndAt   time.Time `query:"endAt"`
}
