package sales

import (
	"context"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/products"
)

type Service struct {
	products products.UseCase
	repo     Repository
}

func NewService(productSvc products.UseCase, r Repository) *Service {
	return &Service{
		products: productSvc,
		repo:     r,
	}
}

func (s *Service) RegisterSale(
	ctx context.Context, desc string, payment PaymentType, cart Cart,
) (Sale, error) {
	var total float64

	items := make([]Item, len(cart.Items))

	for i, item := range cart.Items {
		p, err := s.products.Get(ctx, item.ItemID)
		if err != nil {
			return Sale{}, err
		}

		unitPrice := p.GetUnitPrice(item.Amount)
		total += (unitPrice * float64(item.Amount))
		items[i] = Item{
			Name:      p.Name,
			UnitPrice: unitPrice,
			Amount:    item.Amount,
		}
	}

	sale, err := NewSale(payment, desc, total, items)
	if err != nil {
		return Sale{}, err
	}

	if err := s.repo.Create(ctx, sale); err != nil {
		return Sale{}, err
	}

	return sale, nil
}

func (s *Service) GetAll(ctx context.Context) ([]Sale, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetByPeriod(ctx context.Context, start, end time.Time) ([]Sale, error) {
	return s.repo.Search(ctx, start, end)
}
