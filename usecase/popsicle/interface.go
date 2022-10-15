package popsicle

import (
	"context"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

type Reader interface {
	Get(ctx context.Context, id entity.ID) (entity.Popsicle, error)
	GetAll(ctx context.Context) ([]entity.Popsicle, error)
}

type Writer interface {
	Create(ctx context.Context, p entity.Popsicle) error
	Update(ctx context.Context, p *entity.Popsicle) error
	Delete(ctx context.Context, id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Store(ctx context.Context, flavor string, price float32) (entity.Popsicle, error)
	Get(ctx context.Context, id entity.ID) (entity.Popsicle, error)
	Index(ctx context.Context) ([]entity.Popsicle, error)
	Update(ctx context.Context, p *entity.Popsicle) error
	Delete(ctx context.Context, id entity.ID) error
}
