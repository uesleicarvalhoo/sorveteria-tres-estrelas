package products

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
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
) (entity.Product, error) {
	p, err := entity.NewProduct(name, priceVarejo, priceAtacado, atacadoAmount)
	if err != nil {
		return entity.Product{}, err
	}

	if err := s.r.Create(ctx, p); err != nil {
		return entity.Product{}, err
	}

	return p, nil
}

func (s Service) Get(ctx context.Context, id uuid.UUID) (entity.Product, error) {
	return s.r.Get(ctx, id)
}

func (s Service) Index(ctx context.Context) ([]entity.Product, error) {
	return s.r.GetAll(ctx)
}

func (s Service) Update(ctx context.Context, p *entity.Product) error {
	if err := p.Validate(); err != nil {
		return err
	}

	return s.r.Update(ctx, p)
}

func (s Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.r.Delete(ctx, id)
}
