package user

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

func (s *Service) Create(
	ctx context.Context, name, email, password string, roles ...entity.Permission,
) (entity.User, error) {
	u, err := entity.NewUser(name, email, password, roles...)
	if err != nil {
		return entity.User{}, err
	}

	if err := s.r.Create(ctx, u); err != nil {
		return entity.User{}, err
	}

	return u, nil
}

func (s *Service) Get(ctx context.Context, id uuid.UUID) (entity.User, error) {
	u, err := s.r.Get(ctx, id)
	if err != nil {
		return entity.User{}, err
	}

	return u, nil
}

func (s *Service) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	u, err := s.r.GetByEmail(ctx, email)
	if err != nil {
		return entity.User{}, err
	}

	return u, nil
}
