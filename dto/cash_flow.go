package dto

import (
	"time"
)

type GetCashFlowByPeriodQuery struct {
	StartAt time.Time `query:"startAt"`
	EndAt   time.Time `query:"endAt"`
}
