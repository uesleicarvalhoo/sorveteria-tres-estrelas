package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/transaction"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/transaction/postgres"
	"gorm.io/gorm"
)

func NewTransactionService(db *gorm.DB) transaction.UseCase {
	r := postgres.NewRepository(db)

	return transaction.NewService(r)
}
