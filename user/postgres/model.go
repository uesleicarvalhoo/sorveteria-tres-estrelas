package postgres

import (
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user"
)

type UserModel struct {
	ID           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
}

func (u UserModel) TableName() string { return "users" }

func userToModel(u user.User) UserModel {
	return UserModel{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	}
}

func userModelToEntity(u UserModel) user.User {
	return user.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	}
}
