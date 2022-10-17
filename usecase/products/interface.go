package products

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

type Reader interface {
	Get(ctx context.Context, id uuid.UUID) (entity.Product, error)
	GetAll(ctx context.Context) ([]entity.Product, error)
}

type Writer interface {
	Create(ctx context.Context, p entity.Product) error
	Update(ctx context.Context, p *entity.Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Store(ctx context.Context, name string, varejoPrice, atacadoPrice float64, atacadoAmount int) (entity.Product, error)
	Get(ctx context.Context, id uuid.UUID) (entity.Product, error)
	Index(ctx context.Context) ([]entity.Product, error)
	Update(ctx context.Context, p *entity.Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}
