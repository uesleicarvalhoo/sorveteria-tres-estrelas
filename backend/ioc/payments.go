package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/payment"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/payment/postgres"
	"gorm.io/gorm"
)

func NewPaymentService(db *gorm.DB) payment.UseCase {
	r := postgres.NewRepository(db)

	return payment.NewService(r)
}
