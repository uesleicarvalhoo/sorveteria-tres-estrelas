package sales

import (
	"context"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

type Reader interface {
	Search(ctx context.Context, start, end time.Time) ([]entity.Sale, error)
	GetAll(ctx context.Context) ([]entity.Sale, error)
}

type Writer interface {
	Create(ctx context.Context, s entity.Sale) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	RegisterSale(ctx context.Context, desc string, payment entity.PaymentType, cart entity.Cart) (entity.Sale, error)
	GetAll(ctx context.Context) ([]entity.Sale, error)
	GetByPeriod(ctx context.Context, start, end time.Time) ([]entity.Sale, error)
}
