package sales

import (
	"context"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/products"
)

type Service struct {
	products products.Reader
	repo     Repository
}

func NewService(productsR products.Reader, r Repository) *Service {
	return &Service{
		products: productsR,
		repo:     r,
	}
}

func (s *Service) RegisterSale(
	ctx context.Context, desc string, payment entity.PaymentType, cart entity.Cart,
) (entity.Sale, error) {
	var total float64

	items := make([]entity.SaleItem, len(cart.Items))

	for i, item := range cart.Items {
		p, err := s.products.Get(ctx, item.ItemID)
		if err != nil {
			return entity.Sale{}, err
		}

		unitPrice := p.GetUnitPrice(item.Amount)
		total += (unitPrice * float64(item.Amount))
		items[i] = entity.SaleItem{
			Name:      p.Name,
			UnitPrice: unitPrice,
			Amount:    item.Amount,
		}
	}

	sale, err := entity.NewSale(payment, desc, total, items)
	if err != nil {
		return entity.Sale{}, err
	}

	if err := s.repo.Create(ctx, sale); err != nil {
		return entity.Sale{}, err
	}

	return sale, nil
}

func (s *Service) GetAll(ctx context.Context) ([]entity.Sale, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetByPeriod(ctx context.Context, start, end time.Time) ([]entity.Sale, error) {
	return s.repo.Search(ctx, start, end)
}
