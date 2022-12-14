package payments

import (
	"context"
	"sort"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	r Repository
}

func NewService(r Repository) Service {
	return Service{
		r: r,
	}
}

func (s Service) GetAll(ctx context.Context) ([]Payment, error) {
	payments, err := s.r.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	s.sort(payments)

	return payments, nil
}

func (s Service) GetByPeriod(ctx context.Context, startAt, endAt time.Time) ([]Payment, error) {
	payments, err := s.r.GetBetween(ctx, startAt, endAt)
	if err != nil {
		return nil, err
	}

	s.sort(payments)

	return payments, nil
}

func (s Service) RegisterPayment(ctx context.Context, value float32, desc string) (Payment, error) {
	p := Payment{
		ID:          uuid.New(),
		Value:       value,
		Description: desc,
		CreatedAt:   time.Now(),
	}

	err := s.r.Create(ctx, p)
	if err != nil {
		return Payment{}, err
	}

	return p, nil
}

func (s Service) DeletePayment(ctx context.Context, id uuid.UUID) error {
	return s.r.Delete(ctx, id)
}

func (s Service) UpdatePayment(ctx context.Context, id uuid.UUID, value float32, desc string) (Payment, error) {
	p, err := s.r.Get(ctx, id)
	if err != nil {
		return Payment{}, err
	}

	if value != 0 {
		p.Value = value
	}

	if desc != "" {
		p.Description = desc
	}

	if err := s.r.Update(ctx, &p); err != nil {
		return Payment{}, err
	}

	return p, nil
}

func (s Service) sort(payments []Payment) {
	sort.Slice(payments, func(i, j int) bool {
		return payments[i].CreatedAt.After(payments[j].CreatedAt)
	})
}
