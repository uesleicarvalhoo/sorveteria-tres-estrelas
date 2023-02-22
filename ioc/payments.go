package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payments"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payments/postgres"
	"gorm.io/gorm"
)

func NewPaymentService(db *gorm.DB) payments.UseCase {
	r := postgres.NewRepository(db)

	return payments.NewService(r)
}
