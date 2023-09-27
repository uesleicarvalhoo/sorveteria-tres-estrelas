package transaction

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Reader interface {
	GetAll(ctx context.Context) ([]Transaction, error)
	GetBetween(ctx context.Context, startAt, endAt time.Time) ([]Transaction, error)
}

type Writer interface {
	Create(ctx context.Context, transaction Transaction) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetTransactions(ctx context.Context) ([]Transaction, error)
	GetByPeriod(ctx context.Context, startAt, endAt time.Time) ([]Transaction, error)
	RegisterTransaction(ctx context.Context, value float32, operation Type, desc string) (Transaction, error)
	DeleteTransaction(ctx context.Context, id uuid.UUID) error
}
