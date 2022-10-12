package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r PostgresRepository) Get(ctx context.Context, id uuid.UUID) (sales.Sale, error) {
	var m SaleModel

	if tx := r.db.WithContext(ctx).Preload("Items").First(&m); tx.Error != nil {
		return sales.Sale{}, tx.Error
	}

	return toEntity(m), nil
}

func (r PostgresRepository) GetAll(ctx context.Context) ([]sales.Sale, error) {
	var models []SaleModel

	if tx := r.db.WithContext(ctx).Preload("Items").Find(&models); tx.Error != nil {
		return nil, tx.Error
	}

	sales := make([]sales.Sale, len(models))
	for i, m := range models {
		sales[i] = toEntity(m)
	}

	return sales, nil
}

func (r PostgresRepository) Create(ctx context.Context, s sales.Sale) error {
	m := toModel(s)

	return r.db.WithContext(ctx).Create(&m).Error
}

func (r PostgresRepository) Update(ctx context.Context, s *sales.Sale) error {
	if tx := r.db.WithContext(ctx).Delete(&SaleItemModel{}, "sale_id = ?", s.ID); tx.Error != nil {
		return tx.Error
	}

	return r.db.WithContext(ctx).Save(s).Error
}

func (r PostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&sales.Sale{}, "id = ?", id).Error
}
