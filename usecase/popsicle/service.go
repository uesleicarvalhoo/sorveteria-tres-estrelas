package popsicle

import (
	"context"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
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

func (s Service) Store(ctx context.Context, flavor string, price float64) (entity.Popsicle, error) {
	pop := entity.Popsicle{
		ID:     entity.NewID(),
		Flavor: flavor,
		Price:  price,
	}

	if err := validator.Validate(pop); err != nil {
		return entity.Popsicle{}, err
	}

	if err := s.r.Create(ctx, pop); err != nil {
		return entity.Popsicle{}, err
	}

	return pop, nil
}

func (s Service) Get(ctx context.Context, id entity.ID) (entity.Popsicle, error) {
	return s.r.Get(ctx, id)
}

func (s Service) Index(ctx context.Context) ([]entity.Popsicle, error) {
	return s.r.GetAll(ctx)
}

func (s Service) Update(ctx context.Context, p *entity.Popsicle) error {
	if err := validator.Validate(p); err != nil {
		return err
	}

	return s.r.Update(ctx, p)
}

func (s Service) Delete(ctx context.Context, id entity.ID) error {
	return s.r.Delete(ctx, id)
}
