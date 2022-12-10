package dto

import "time"

type GetPaymentByPeriodQuery struct {
	StartAt time.Time `query:"startAt"`
	EndAt   time.Time `query:"endAt"`
}

type CreatePaymentPayload struct {
	Value       float32 `json:"value"`
	Description string  `json:"description"`
}

type UpdatePaymentPayload struct {
	ID          string  `json:"id"`
	Value       float32 `json:"value"`
	Description string  `json:"description"`
}
