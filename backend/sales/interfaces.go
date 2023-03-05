package sales

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Reader interface {
	Search(ctx context.Context, start, end time.Time) ([]Sale, error)
	GetAll(ctx context.Context) ([]Sale, error)
}

type Writer interface {
	Create(ctx context.Context, s Sale) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	RegisterSale(ctx context.Context, desc string, payment PaymentType, cart Cart) (Sale, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]Sale, error)
	GetByPeriod(ctx context.Context, start, end time.Time) ([]Sale, error)
}
