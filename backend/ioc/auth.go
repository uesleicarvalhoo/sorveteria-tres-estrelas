package ioc

import (
	"github.com/kong/go-kong/kong"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/auth"
	"gorm.io/gorm"
)

func NewAuthService(
	db *gorm.DB, kongCli *kong.Client, secret, kongConsumer, kongJwtKey string,
) auth.UseCase {
	userSvc := NewUserService(db)
	provider := auth.NewKongProvider(kongCli, kongConsumer, kongJwtKey)

	return auth.NewService(userSvc, provider)
}
