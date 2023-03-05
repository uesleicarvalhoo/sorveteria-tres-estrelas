package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/cashflow"
	"gorm.io/gorm"
)

func NewCashFlowService(db *gorm.DB) cashflow.UseCase {
	paymentSvc := NewPaymentService(db)
	saleSvc := NewSaleService(db)

	return cashflow.NewService(paymentSvc, saleSvc)
}
