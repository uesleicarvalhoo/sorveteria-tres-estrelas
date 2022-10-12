package sales

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/pkg/validator"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/popsicle"
)

type Reader interface {
	GetByPeriod(ctx context.Context, start, end time.Time) ([]Sale, error)
}

type Writer interface {
	Create(ctx context.Context, s Sale) error
}

type Repository interface {
	Reader
	Writer
}

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

func (s *Service) NewSale(ctx context.Context, description string, paymentT PaymentType, cart Cart) (Sale, error) {
	var total float32

	items := make([]SaleItem, len(cart.Items))

	for i, item := range cart.Items {
		p, err := s.popsicles.Get(ctx, item.PopsicleID)
		if err != nil {
			return Sale{}, err
		}

		total += (p.Price * float32(item.Amount))
		items[i] = SaleItem{
			Name:   fmt.Sprintf("Picole de %s", p.Flavor),
			Price:  p.Price,
			Amount: item.Amount,
		}
	}

	sale := Sale{
		ID:          uuid.New(),
		Total:       total,
		Items:       items,
		Description: description,
		Date:        time.Now(),
	}

	if err := validator.Validate(sale); err != nil {
		return Sale{}, err
	}

	if err := s.repo.Create(ctx, sale); err != nil {
		return Sale{}, err
	}

	return sale, nil
}

func (s *Service) GetByPeriod(ctx context.Context, start, end time.Time) ([]Sale, error) {
	return s.repo.GetByPeriod(ctx, start, end)
}
