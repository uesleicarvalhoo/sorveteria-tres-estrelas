package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances"
	"gorm.io/gorm"
)

type BalancesPostgres struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *BalancesPostgres {
	return &BalancesPostgres{
		db: db,
	}
}

func (r BalancesPostgres) Get(ctx context.Context, id uuid.UUID) (balances.Balance, error) {
	var b balances.Balance

	if tx := r.db.WithContext(ctx).First(&b, "id = ?", id); tx.Error != nil {
		return balances.Balance{}, tx.Error
	}

	return b, nil
}

func (r BalancesPostgres) GetAll(ctx context.Context) ([]balances.Balance, error) {
	var records []balances.Balance

	if tx := r.db.WithContext(ctx).Find(&records); tx.Error != nil {
		return nil, tx.Error
	}

	return records, nil
}

func (r BalancesPostgres) GetBetween(ctx context.Context, startAt, endAt time.Time) ([]balances.Balance, error) {
	var records []balances.Balance
	if tx := r.db.WithContext(ctx).Find(&records, "created_at BETWEEN ? AND ?", startAt, endAt); tx.Error != nil {
		return nil, tx.Error
	}

	return records, nil
}

func (r BalancesPostgres) Create(ctx context.Context, b balances.Balance) error {
	return r.db.WithContext(ctx).Create(&b).Error
}

func (r BalancesPostgres) Update(ctx context.Context, b *balances.Balance) error {
	return r.db.WithContext(ctx).Save(b).Error
}

func (r BalancesPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&balances.Balance{}, "id = ?", id).Error
}
