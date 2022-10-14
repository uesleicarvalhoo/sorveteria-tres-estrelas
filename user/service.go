package user

import (
	"context"

	"github.com/google/uuid"
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

func (s *Service) Create(ctx context.Context, name, email, passwdHash string, roles ...Permission) (User, error) {
	u := User{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: passwdHash,
		Permissions:  roles,
	}

	if err := validator.Validate(u); err != nil {
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
