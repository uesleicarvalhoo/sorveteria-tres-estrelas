package popsicle

import (
	"context"

	"github.com/google/uuid"
)

type Popsicle struct {
	ID     uuid.UUID `json:"id"`
	Flavor string    `json:"flavor" validate:"required,min=4"`
	Price  float32   `json:"price" validate:"required"`
}

type Reader interface {
	Get(ctx context.Context, id uuid.UUID) (Popsicle, error)
	GetAll(ctx context.Context) ([]Popsicle, error)
}

type Writer interface {
	Create(ctx context.Context, p Popsicle) error
	Update(ctx context.Context, p *Popsicle) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Repository interface {
	Reader
	Writer
}
