package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/cashflow"
	"gorm.io/gorm"
)

func NewCashFlowService(db *gorm.DB) cashflow.UseCase {
	saleSvc := NewSaleService(db)
	transactionSvc := NewTransactionService(db)

	return cashflow.NewService(saleSvc, transactionSvc)
}
