package jwt

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user"
)

type Service struct {
	secret string
}

func NewService(secretKey string) Service {
	return Service{
		secret: secretKey,
	}
}

func (s Service) Generate(ctx context.Context, u user.User, exp time.Time) (string, error) {
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"kid":   "guciv3GWY8iCWcmromU2olQsQf5VeLIA",
			"sub":   u.ID.String(),
			"name":  u.Name,
			"email": u.Email,
			"iat":   time.Now().Unix(),
			"exp":   exp.Unix(),
		})

	return claims.SignedString([]byte(s.secret))
}

func (s Service) Validate(ctx context.Context, token string) (user.User, error) {
	return validate(ctx, token, s.secret)
}
