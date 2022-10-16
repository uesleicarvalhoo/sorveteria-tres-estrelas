package dto

import "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"

type CreateUserPayload struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Permissions []entity.Permission
}
