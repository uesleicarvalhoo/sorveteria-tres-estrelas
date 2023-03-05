package jwt

import (
	"context"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user"
)

type UseCase interface {
	Generate(ctx context.Context, u user.User, exp time.Time) (string, error)
	Validate(ctx context.Context, token string) (user.User, error)
}
