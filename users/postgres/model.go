package postgres

import (
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
)

type UserModel struct {
	ID           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
	Permissions  string
}

func (u UserModel) TableName() string { return "users" }

func userToModel(u users.User) UserModel {
	return UserModel{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	}
}

func userModelToEntity(u UserModel) users.User {
	return users.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	}
}
