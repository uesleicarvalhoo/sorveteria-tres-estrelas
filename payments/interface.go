package payments

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Reader interface {
	Get(ctx context.Context, id uuid.UUID) (Payment, error)
	GetAll(ctx context.Context) ([]Payment, error)
	GetBetween(ctx context.Context, startAt, endAt time.Time) ([]Payment, error)
}

type Writer interface {
	Update(ctx context.Context, payment *Payment) error
	Create(ctx context.Context, payment Payment) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	RegisterPayment(ctx context.Context, value float32, desc string) (Payment, error)
	DeletePayment(ctx context.Context, id uuid.UUID) error
	GetByPeriod(ctx context.Context, startAt, endAt time.Time) ([]Payment, error)
	GetAll(ctx context.Context) ([]Payment, error)
	UpdatePayment(ctx context.Context, id uuid.UUID, value float32, desc string) (Payment, error)
}
