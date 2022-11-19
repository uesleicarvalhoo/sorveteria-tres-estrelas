package auth

import (
	"context"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
)

type UseCase interface {
	Login(ctx context.Context, email, password string) (JwtToken, error)
	RefreshToken(ctx context.Context, token string) (JwtToken, error)
	Authorize(ctx context.Context, token string) (users.User, error)
}
