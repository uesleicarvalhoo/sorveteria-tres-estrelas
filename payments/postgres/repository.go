package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payments"
	"gorm.io/gorm"
)

type PaymentsPostgres struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *PaymentsPostgres {
	return &PaymentsPostgres{
		db: db,
	}
}

func (r PaymentsPostgres) Get(ctx context.Context, id uuid.UUID) (payments.Payment, error) {
	var p payments.Payment

	if tx := r.db.WithContext(ctx).First(&p, "id = ?", id); tx.Error != nil {
		return payments.Payment{}, tx.Error
	}

	return p, nil
}

func (r PaymentsPostgres) GetAll(ctx context.Context) ([]payments.Payment, error) {
	var records []payments.Payment

	if tx := r.db.WithContext(ctx).Find(&records); tx.Error != nil {
		return nil, tx.Error
	}

	return records, nil
}

func (r PaymentsPostgres) GetBetween(ctx context.Context, startAt, endAt time.Time) ([]payments.Payment, error) {
	var records []payments.Payment
	if tx := r.db.WithContext(ctx).Find(&records, "created_at BETWEEN ? AND ?", startAt, endAt); tx.Error != nil {
		return nil, tx.Error
	}

	return records, nil
}

func (r PaymentsPostgres) Create(ctx context.Context, p payments.Payment) error {
	return r.db.WithContext(ctx).Create(&p).Error
}

func (r PaymentsPostgres) Update(ctx context.Context, p *payments.Payment) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r PaymentsPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&payments.Payment{}, "id = ?", id).Error
}
