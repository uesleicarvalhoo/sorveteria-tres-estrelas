package repository

import (
	"strings"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

const sep = "|"

type UserModel struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
	Permissions  string
}

func (u UserModel) TableName() string { return "users" }

func userToModel(u entity.User) UserModel {
	perm := make([]string, len(u.Permissions))
	for i, role := range u.Permissions {
		perm[i] = string(role)
	}

	return UserModel{
		ID:           u.ID.String(),
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Permissions:  strings.Join(perm, sep),
	}
}

func userModelToEntity(u UserModel) entity.User {
	splited := strings.Split(u.Permissions, sep)
	permissions := make([]entity.Permission, len(splited))

	for i, p := range splited {
		permissions[i] = entity.Permission(p)
	}

	id, _ := entity.StringToID(u.ID)

	return entity.User{
		ID:           id,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Permissions:  permissions,
	}
}
