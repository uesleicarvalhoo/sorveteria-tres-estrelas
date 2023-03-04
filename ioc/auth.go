package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cache"
	"gorm.io/gorm"
)

func NewAuthService(secretKey string, db *gorm.DB, cache cache.Cache) auth.UseCase {
	userSvc := NewUserService(db)

	return auth.NewService(secretKey, userSvc, cache)
}
