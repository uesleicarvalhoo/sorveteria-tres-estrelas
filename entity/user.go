package entity

import (
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/pkg/validator"
	"golang.org/x/crypto/bcrypt"
)

type Permission string

const (
	ReadWriteSalesRole = "sales:read,write"
	ReadSalesRole      = "sales:read"

	ReadWritePopsicle = "popsicle:read,write"
	ReadPopsicle      = "popsicle:read"
)

const minPasswordLength = 5

type User struct {
	ID           uuid.UUID    `json:"id"`
	Name         string       `json:"name" validate:"required"`
	Email        string       `json:"email" validate:"email"`
	PasswordHash string       `json:"-"`
	Permissions  []Permission `json:"roles"`
}

func NewUser(name, email, password string, permissions ...Permission) (User, error) {
	pwd, err := generatePassword(password)
	if err != nil {
		return User{}, err
	}

	u := User{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: pwd,
		Permissions:  permissions,
	}

	if err := u.Validate(); err != nil {
		return User{}, err
	}

	return u, nil
}

func generatePassword(password string) (string, error) {
	if len(password) < minPasswordLength {
		return "", ErrTooShortPassword
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

func (u User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	return err == nil
}

func (u User) Validate() error {
	return validator.Validate(u)
}
