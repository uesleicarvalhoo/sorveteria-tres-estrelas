package repository

import (
	"context"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r UserPostgres) Get(ctx context.Context, id entity.ID) (entity.User, error) {
	var m UserModel

	if tx := r.db.WithContext(ctx).First(&m, "id = ?", id); tx.Error != nil {
		return entity.User{}, tx.Error
	}

	return userModelToEntity(m), nil
}

func (r UserPostgres) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var m UserModel

	if tx := r.db.WithContext(ctx).First(&m, "email = ?", email); tx.Error != nil {
		return entity.User{}, tx.Error
	}

	return userModelToEntity(m), nil
}

func (r UserPostgres) GetAll(ctx context.Context) ([]entity.User, error) {
	var records []UserModel

	if tx := r.db.WithContext(ctx).Find(&records); tx.Error != nil {
		return nil, tx.Error
	}

	users := make([]entity.User, len(records))
	for i, model := range records {
		users[i] = userModelToEntity(model)
	}

	return users, nil
}

func (r UserPostgres) Create(ctx context.Context, u entity.User) error {
	m := userToModel(u)

	return r.db.WithContext(ctx).Create(&m).Error
}

func (r UserPostgres) Update(ctx context.Context, u entity.User) error {
	m := userToModel(u)

	return r.db.WithContext(ctx).Save(&m).Error
}

func (r UserPostgres) Delete(ctx context.Context, id entity.ID) error {
	return r.db.WithContext(ctx).Delete(&UserModel{}, "id = ?", id).Error
}
