package jwt

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kong/go-kong/kong"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user"
)

type Kong struct {
	key      string
	username string
	cli      *kong.Client
}

func NewKongService(cli *kong.Client, username, key string) Kong {
	return Kong{
		key:      key,
		username: username,
		cli:      cli,
	}
}

func (k Kong) getConsumer(ctx context.Context, username string) (*kong.Consumer, error) {
	return k.cli.Consumers.Get(ctx, &username)
}

func (k Kong) getAuthJwt(ctx context.Context) (*kong.JWTAuth, error) {
	consumer, err := k.getConsumer(ctx, k.username)
	if err != nil {
		return nil, err
	}

	auth, err := k.cli.JWTAuths.Get(ctx, consumer.ID, &k.key)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (k Kong) Generate(ctx context.Context, u user.User, exp time.Time) (string, error) {
	auth, err := k.getAuthJwt(ctx)
	if err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"iss":   *auth.Key,
			"sub":   u.ID.String(),
			"name":  u.Name,
			"email": u.Email,
			"iat":   time.Now().Unix(),
			"exp":   exp.Unix(),
		})

	return claims.SignedString([]byte(*auth.Secret))
}

func (k Kong) Validate(ctx context.Context, token string) (user.User, error) {
	auth, err := k.getAuthJwt(ctx)
	if err != nil {
		return user.User{}, err
	}

	return validate(ctx, token, *auth.Key)
}
