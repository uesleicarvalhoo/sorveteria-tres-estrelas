package auth

import (
	"context"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/user"
)

type UseCase interface {
	Login(ctx context.Context, payload LoginPayload) (JwtToken, error)
	RefreshToken(ctx context.Context, payload RefreshTokenPayload) (JwtToken, error)
	Authorize(ctx context.Context, token string) (user.User, error)
}

type ConfigProvider interface {
	GetSecretKey(ctx context.Context) (string, error)
	GetIssuer(ctx context.Context) (string, error)
}
