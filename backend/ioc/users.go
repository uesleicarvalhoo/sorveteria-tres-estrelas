package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/user"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/user/postgres"
	"gorm.io/gorm"
)

func NewUserService(db *gorm.DB) user.UseCase {
	return user.NewService(postgres.NewRepository(db))
}
