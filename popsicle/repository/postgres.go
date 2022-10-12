package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/popsicle"
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

func (r PostgresRepository) Get(ctx context.Context, id uuid.UUID) (popsicle.Popsicle, error) {
	var p popsicle.Popsicle

	if tx := r.db.WithContext(ctx).First(&p); tx.Error != nil {
		return popsicle.Popsicle{}, tx.Error
	}

	return p, nil
}

func (r PostgresRepository) GetAll(ctx context.Context) ([]popsicle.Popsicle, error) {
	var res []popsicle.Popsicle

	if tx := r.db.WithContext(ctx).Find(&res); tx.Error != nil {
		return nil, tx.Error
	}

	return res, nil
}

func (r PostgresRepository) Create(ctx context.Context, p popsicle.Popsicle) error {
	return r.db.WithContext(ctx).Create(&p).Error
}

func (r PostgresRepository) Update(ctx context.Context, p *popsicle.Popsicle) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r PostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&popsicle.Popsicle{}, "id = ?", id).Error
}
