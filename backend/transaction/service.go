package transaction

import (
	"context"
	"sort"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	r Repository
}

func NewService(repo Repository) UseCase {
	return &Service{
		r: repo,
	}
}

func (s *Service) GetTransactions(ctx context.Context) ([]Transaction, error) {
	tr, err := s.r.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	s.sort(tr)

	return tr, nil
}

func (s *Service) GetByPeriod(ctx context.Context, startAt, endAt time.Time) ([]Transaction, error) {
	tr, err := s.r.GetBetween(ctx, startAt, endAt)
	if err != nil {
		return nil, err
	}

	s.sort(tr)

	return tr, nil
}

func (s *Service) RegisterTransaction(
	ctx context.Context, value float32, operation Type, desc string,
) (Transaction, error) {
	if err := operation.Validate(); err != nil {
		return Transaction{}, err
	}

	tr := Transaction{
		ID:          uuid.New(),
		Value:       value,
		Description: desc,
		Type:        operation,
		CreatedAt:   time.Now(),
	}

	if err := s.r.Create(ctx, tr); err != nil {
		return Transaction{}, err
	}

	return tr, nil
}

func (s *Service) DeleteTransaction(ctx context.Context, id uuid.UUID) error {
	return s.r.Delete(ctx, id)
}

func (s Service) sort(transactions []Transaction) {
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].CreatedAt.After(transactions[j].CreatedAt)
	})
}
