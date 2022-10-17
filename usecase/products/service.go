package products

import (
	"context"

	"github.com/google/uuid"
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

func (s Service) Store(
	ctx context.Context, name string, varejoPrice, atacadoPrice float64, atacadoMinAmount int,
) (entity.Product, error) {
	p := entity.Product{
		ID:            uuid.New(),
		Name:          name,
		PriceVarejo:   varejoPrice,
		PriceAtacado:  atacadoPrice,
		AtacadoAmount: atacadoMinAmount,
	}

	if err := validator.Validate(p); err != nil {
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
	if err := validator.Validate(p); err != nil {
		return err
	}

	return s.r.Update(ctx, p)
}

func (s Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.r.Delete(ctx, id)
}
