package balances

import (
	"context"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
)

type Reader interface {
	GetAll(ctx context.Context) ([]Balance, error)
	GetBetween(ctx context.Context, start, end time.Time) ([]Balance, error)
}

type Writer interface {
	Create(ctx context.Context, b Balance) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	RegisterOperation(ctx context.Context, value float32, desc string, tp OperationType) (Balance, error)
	RegisterFromSale(ctx context.Context, sale sales.Sale) (Balance, error)
	GetAll(ctx context.Context) ([]Balance, error)
	GetCashFlow(ctx context.Context) (CashFlow, error)
	GetCashFlowBetween(ctx context.Context, startAt, endAt time.Time) (CashFlow, error)
}
