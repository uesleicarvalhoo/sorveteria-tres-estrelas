package products

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{
		r: r,
	}
}

func (s Service) Store(
	ctx context.Context, name string, priceVarejo, priceAtacado float64, atacadoAmount int,
) (Product, error) {
	p, err := NewProduct(name, priceVarejo, priceAtacado, atacadoAmount)
	if err != nil {
		return Product{}, err
	}

	if err := s.r.Create(ctx, p); err != nil {
		return Product{}, err
	}

	return p, nil
}

func (s Service) Get(ctx context.Context, id uuid.UUID) (Product, error) {
	return s.r.Get(ctx, id)
}

func (s Service) Index(ctx context.Context) ([]Product, error) {
	return s.r.GetAll(ctx)
}

func (s Service) Update(ctx context.Context, id uuid.UUID, payload UpdatePayload) (Product, error) {
	if payload.IsEmpty() {
		return Product{}, ErrNoDataForUpdate
	}

	p, err := s.r.Get(ctx, id)
	if err != nil {
		return Product{}, err
	}

	if payload.Name != "" {
		p.Name = payload.Name
	}

	if payload.PriceVarejo != 0 {
		p.PriceVarejo = payload.PriceVarejo
	}

	if payload.PriceAtacado != 0 {
		p.PriceAtacado = payload.PriceAtacado
	}

	if payload.AtacadoAmount != 0 {
		p.AtacadoAmount = payload.AtacadoAmount
	}

	if err := p.Validate(); err != nil {
		return Product{}, err
	}

	if err := s.r.Update(ctx, &p); err != nil {
		return Product{}, err
	}

	return p, nil
}

func (s Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.r.Delete(ctx, id)
}
