package product

import (
	"context"

	"github.com/google/uuid"
)

type Reader interface {
	Get(ctx context.Context, id uuid.UUID) (Product, error)
	GetAll(ctx context.Context) ([]Product, error)
}

type Writer interface {
	Create(ctx context.Context, p Product) error
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Store(ctx context.Context, name string, varejoPrice, atacadoPrice float64, atacadoAmount int) (Product, error)
	Get(ctx context.Context, id uuid.UUID) (Product, error)
	Index(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, id uuid.UUID, payload UpdatePayload) (Product, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
