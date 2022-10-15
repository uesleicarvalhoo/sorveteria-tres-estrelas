package sales

import (
	"context"
	"fmt"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/pkg/validator"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/popsicle"
)

type Service struct {
	popsicles popsicle.Reader
	repo      Repository
}

func NewService(popsicleR popsicle.Reader, r Repository) *Service {
	return &Service{
		popsicles: popsicleR,
		repo:      r,
	}
}

func (s *Service) NewSale(
	ctx context.Context, description string, payment entity.PaymentType, cart entity.Cart,
) (entity.Sale, error) {
	var total float64

	items := make([]entity.SaleItem, len(cart.Items))

	for i, item := range cart.Items {
		p, err := s.popsicles.Get(ctx, item.PopsicleID)
		if err != nil {
			return entity.Sale{}, err
		}

		total += (p.Price * float64(item.Amount))
		items[i] = entity.SaleItem{
			Name:   fmt.Sprintf("Picole de %s", p.Flavor),
			Price:  p.Price,
			Amount: item.Amount,
		}
	}

	sale := entity.Sale{
		ID:          entity.NewID(),
		Total:       total,
		Items:       items,
		Description: description,
		Date:        time.Now(),
	}

	if err := validator.Validate(sale); err != nil {
		return entity.Sale{}, err
	}

	if err := s.repo.Create(ctx, sale); err != nil {
		return entity.Sale{}, err
	}

	return sale, nil
}

func (s *Service) GetByPeriod(ctx context.Context, start, end time.Time) ([]entity.Sale, error) {
	return s.repo.Search(ctx, start, end)
}
