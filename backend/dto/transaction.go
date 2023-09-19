package dto

import "time"

type GetTransactionByPeriodQuery struct {
	StartAt time.Time `query:"startAt"`
	EndAt   time.Time `query:"endAt"`
}

type CreateTransactionPayload struct {
	Value       float32 `json:"value"`
	Description string  `json:"description"`
}
