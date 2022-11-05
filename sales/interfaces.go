package sales

import (
	"context"
	"time"
)

type Reader interface {
	Search(ctx context.Context, start, end time.Time) ([]Sale, error)
	GetAll(ctx context.Context) ([]Sale, error)
}

type Writer interface {
	Create(ctx context.Context, s Sale) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	RegisterSale(ctx context.Context, desc string, payment PaymentType, cart Cart) (Sale, error)
	GetAll(ctx context.Context) ([]Sale, error)
	GetByPeriod(ctx context.Context, start, end time.Time) ([]Sale, error)
}
