package popsicle

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/pkg/validator"
)

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{
		r: r,
	}
}

func (s Service) Create(ctx context.Context, flavor string, price float32) (Popsicle, error) {
	pop := Popsicle{
		ID:     uuid.New(),
		Flavor: flavor,
		Price:  price,
	}

	if err := validator.Validate(pop); err != nil {
		return Popsicle{}, err
	}

	if err := s.r.Create(ctx, pop); err != nil {
		return Popsicle{}, err
	}

	return pop, nil
}

func (s Service) Get(ctx context.Context, id uuid.UUID) (Popsicle, error) {
	return s.r.Get(ctx, id)
}

func (s Service) GetAll(ctx context.Context) ([]Popsicle, error) {
	return s.r.GetAll(ctx)
}

func (s Service) Update(ctx context.Context, p *Popsicle) error {
	if err := validator.Validate(p); err != nil {
		return err
	}

	return s.r.Update(ctx, p)
}

func (s Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.r.Delete(ctx, id)
}
