package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/transaction"
	"gorm.io/gorm"
)

type TransactionsPostgres struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *TransactionsPostgres {
	return &TransactionsPostgres{
		db: db,
	}
}

func (r TransactionsPostgres) Get(ctx context.Context, id uuid.UUID) (transaction.Transaction, error) {
	var p transaction.Transaction

	if tx := r.db.WithContext(ctx).First(&p, "id = ?", id); tx.Error != nil {
		return transaction.Transaction{}, tx.Error
	}

	return p, nil
}

func (r TransactionsPostgres) GetAll(ctx context.Context) ([]transaction.Transaction, error) {
	var records []transaction.Transaction

	if tx := r.db.WithContext(ctx).Find(&records); tx.Error != nil {
		return nil, tx.Error
	}

	return records, nil
}

func (r TransactionsPostgres) GetBetween(
	ctx context.Context, startAt, endAt time.Time,
) ([]transaction.Transaction, error) {
	var records []transaction.Transaction
	if tx := r.db.WithContext(ctx).Find(&records, "created_at BETWEEN ? AND ?", startAt, endAt); tx.Error != nil {
		return nil, tx.Error
	}

	return records, nil
}

func (r TransactionsPostgres) Create(ctx context.Context, p transaction.Transaction) error {
	return r.db.WithContext(ctx).Create(&p).Error
}

func (r TransactionsPostgres) Update(ctx context.Context, p *transaction.Transaction) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r TransactionsPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&transaction.Transaction{}, "id = ?", id).Error
}
