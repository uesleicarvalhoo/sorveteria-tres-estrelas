package postgres

import (
	"strings"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
)

const sep = "|"

type UserModel struct {
	ID           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
	Permissions  string
}

func (u UserModel) TableName() string { return "users" }

func userToModel(u users.User) UserModel {
	perm := make([]string, len(u.Permissions))
	for i, role := range u.Permissions {
		perm[i] = string(role)
	}

	return UserModel{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Permissions:  strings.Join(perm, sep),
	}
}

func userModelToEntity(u UserModel) users.User {
	splited := strings.Split(u.Permissions, sep)
	permissions := make([]users.Permission, len(splited))

	for i, p := range splited {
		permissions[i] = users.Permission(p)
	}

	return users.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Permissions:  permissions,
	}
}
