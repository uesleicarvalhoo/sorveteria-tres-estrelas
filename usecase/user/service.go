package user

import (
	"context"

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

func (s *Service) Create(
	ctx context.Context, name, email, passwdHash string, roles ...entity.Permission,
) (entity.User, error) {
	u := entity.User{
		ID:           entity.NewID(),
		Name:         name,
		Email:        email,
		PasswordHash: passwdHash,
		Permissions:  roles,
	}

	if err := validator.Validate(u); err != nil {
		return entity.User{}, err
	}

	if err := s.r.Create(ctx, u); err != nil {
		return entity.User{}, err
	}

	return u, nil
}

func (s *Service) Get(ctx context.Context, id entity.ID) (entity.User, error) {
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
