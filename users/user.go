package users

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity/validator"
	"golang.org/x/crypto/bcrypt"
)

const minPasswordLength = 5

var ErrTooShortPassword = fmt.Errorf("a senha precisa conter ao menos %d caracters", minPasswordLength)

type User struct {
	ID           uuid.UUID    `json:"id"`
	Name         string       `json:"name"`
	Email        string       `json:"email"`
	PasswordHash string       `json:"-"`
	Permissions  []Permission `json:"permissions"`
}

func NewUser(name, email, password string, permissions ...Permission) (User, error) {
	pwd, err := generatePassword(password)
	if err != nil {
		return User{}, err
	}

	if len(permissions) == 0 {
		permissions = DefaultPermissions()
	}

	u := User{
		ID:           uuid.New(),
		Name:         strings.TrimSpace(name),
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
	v := validator.New()
	if u.Name == "" {
		v.AddError("nome", "campo obrigatório")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		v.AddError("email", "campo invalido")
	}

	return v.Validate()
}

func (u User) AuthorizeDomainAction(domain, action string) bool {
	for _, p := range u.Permissions {
		d, permission := p.getDomainActions()

		if d == domain {
			for _, perm := range permission {
				if perm == action {
					return true
				}
			}
		}
	}

	return false
}

func (u User) HasPermission(p Permission) bool {
	for _, up := range u.Permissions {
		if p.Domain() == up.Domain() && strings.Contains(up.StrActions(), p.StrActions()) {
			return true
		}
	}

	return false
}
