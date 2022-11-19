package dto

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
)

type CreateUserPayload struct {
	Name        string             `json:"name"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	Permissions []users.Permission `json:"permissions"`
}
