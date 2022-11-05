package users

import (
	"context"

	"github.com/google/uuid"
)

type Reader interface {
	Get(ctx context.Context, id uuid.UUID) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
}

type Writer interface {
	Create(ctx context.Context, u User) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Get(ctx context.Context, id uuid.UUID) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Create(ctx context.Context, name, email, password string, permissions ...Permission) (User, error)
}
