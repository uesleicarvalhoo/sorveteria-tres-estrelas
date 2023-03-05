package jwt

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user"
)

func validate(_ context.Context, token, secret string) (user.User, error) {
	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return user.User{}, err
	}

	claims, ok := tokenObj.Claims.(jwt.MapClaims)
	if !ok || !tokenObj.Valid {
		return user.User{}, ErrInvalidToken
	}

	sub, _ := claims["sub"].(string)
	name, _ := claims["name"].(string)
	email, _ := claims["email"].(string)

	id, err := uuid.Parse(sub)
	if err != nil {
		return user.User{}, err
	}

	return user.User{
		ID:    id,
		Name:  name,
		Email: email,
	}, nil
}
