package user

import (
	"context"

	"github.com/google/uuid"
)

type Permission string

const (
	ReadWriteSalesRole = "sales:read,write"
	ReadSalesRole      = "sales:read"

	ReadWritePopsicle = "popsicle:read,write"
	ReadPopsicle      = "popsicle:read"
)

type User struct {
	ID           uuid.UUID    `json:"id"`
	Name         string       `json:"name" validate:"required"`
	Email        string       `json:"email" validate:"email"`
	PasswordHash string       `json:"-"`
	Permissions  []Permission `json:"roles"`
}

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
