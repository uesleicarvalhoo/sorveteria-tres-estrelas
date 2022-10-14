package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user"
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

func (r PostgresRepository) Get(ctx context.Context, id uuid.UUID) (user.User, error) {
	var m UserModel

	if tx := r.db.WithContext(ctx).First(&m, "id = ?", id); tx.Error != nil {
		return user.User{}, tx.Error
	}

	return toEntity(m), nil
}

func (r PostgresRepository) GetByEmail(ctx context.Context, email string) (user.User, error) {
	var m UserModel

	if tx := r.db.WithContext(ctx).First(&m, "email = ?", email); tx.Error != nil {
		return user.User{}, tx.Error
	}

	return toEntity(m), nil
}

func (r PostgresRepository) GetAll(ctx context.Context) ([]user.User, error) {
	var records []UserModel

	if tx := r.db.WithContext(ctx).Find(&records); tx.Error != nil {
		return nil, tx.Error
	}

	users := make([]user.User, len(records))
	for i, model := range records {
		users[i] = toEntity(model)
	}

	return users, nil
}

func (r PostgresRepository) Create(ctx context.Context, u user.User) error {
	m := toModel(u)

	return r.db.WithContext(ctx).Create(&m).Error
}

func (r PostgresRepository) Update(ctx context.Context, u user.User) error {
	m := toModel(u)

	return r.db.WithContext(ctx).Save(&m).Error
}

func (r PostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&UserModel{}, "id = ?", id).Error
}
