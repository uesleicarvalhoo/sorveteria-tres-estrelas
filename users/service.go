package users

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

func (s *Service) Create(ctx context.Context, name, email, password string) (User, error) {
	u, err := NewUser(name, email, password)
	if err != nil {
		return User{}, err
	}

	if err := s.r.Create(ctx, u); err != nil {
		return User{}, err
	}

	return u, nil
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (User, error) {
	u, err := s.r.Get(ctx, id)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (s *Service) GetByEmail(ctx context.Context, email string) (User, error) {
	u, err := s.r.GetByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}

	return u, nil
}
