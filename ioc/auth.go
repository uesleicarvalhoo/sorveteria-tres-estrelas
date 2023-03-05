package ioc

import (
	"github.com/kong/go-kong/kong"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth/jwt"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/infrastructure/cache"
	"gorm.io/gorm"
)

func NewAuthService(
	db *gorm.DB, cache cache.Cache, kongCli *kong.Client, secret, kongConsumer, kongJwtKey string,
) auth.UseCase {
	userSvc := NewUserService(db)

	jwtSvc := jwt.NewKongService(kongCli, kongConsumer, kongJwtKey)

	return auth.NewService(userSvc, jwtSvc, cache)
}
