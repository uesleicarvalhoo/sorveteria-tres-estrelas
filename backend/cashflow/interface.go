package cashflow

import (
	"context"
	"time"
)

type UseCase interface {
	GetCashFlow(ctx context.Context) (CashFlow, error)
	GetCashFlowBetween(ctx context.Context, startAt, endAt time.Time) (CashFlow, error)
}
