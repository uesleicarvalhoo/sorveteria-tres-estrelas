package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payment"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payment/postgres"
	"gorm.io/gorm"
)

func NewPaymentService(db *gorm.DB) payment.UseCase {
	r := postgres.NewRepository(db)

	return payment.NewService(r)
}
