package user

import (
	"context"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

type Reader interface {
	Get(ctx context.Context, id entity.ID) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
}

type Writer interface {
	Create(ctx context.Context, u entity.User) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Get(ctx context.Context, id entity.ID) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	Create(ctx context.Context, name, email, password string, permissions ...entity.Permission) (entity.User, error)
}
