package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/user"
)

func GenerateJwtToken(_ context.Context, u user.User, exp time.Time, issuer, secretKey string) (string, error) {
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"iss":   issuer,
			"sub":   u.ID.String(),
			"name":  u.Name,
			"email": u.Email,
			"iat":   time.Now().Unix(),
			"exp":   exp.Unix(),
		})

	return claims.SignedString([]byte(secretKey))
}

func ValidateJwtToken(_ context.Context, token, secretKey string) (user.User, error) {
	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}

		return []byte(secretKey), nil
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
