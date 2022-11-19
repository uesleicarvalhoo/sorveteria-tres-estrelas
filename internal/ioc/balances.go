package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances/postgres"
	"gorm.io/gorm"
)

func NewBalanceService(db *gorm.DB) balances.UseCase {
	r := postgres.NewRepository(db)

	return balances.NewService(r)
}
