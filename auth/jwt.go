package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

type JwtToken struct {
	GrantType    string `json:"grant_type"`
	AcessToken   string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expiration"`
}

func GenerateJwtToken(secret string, userID entity.ID, exp time.Time) (string, error) {
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.StandardClaims{
			Subject:   userID.String(),
			ExpiresAt: exp.Unix(),
		})

	return claims.SignedString([]byte(secret))
}

func ValidateJwtToken(token, secret string) (entity.ID, error) {
	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return entity.ID{}, err
	}

	claims, ok := tokenObj.Claims.(jwt.MapClaims)
	if ok && tokenObj.Valid {
		sub, _ := claims["sub"].(string)

		return entity.StringToID(sub)
	}

	return entity.ID{}, ErrInvalidToken
}
