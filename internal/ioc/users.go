package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users/postgres"
	"gorm.io/gorm"
)

func NewUserService(db *gorm.DB) users.UseCase {
	return users.NewService(postgres.NewRepository(db))
}
