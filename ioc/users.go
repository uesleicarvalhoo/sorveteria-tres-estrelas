package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user/postgres"
	"gorm.io/gorm"
)

func NewUserService(db *gorm.DB) user.UseCase {
	return user.NewService(postgres.NewRepository(db))
}
