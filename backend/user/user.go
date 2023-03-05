package user

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/validator"
	"golang.org/x/crypto/bcrypt"
)

const minPasswordLength = 5

var ErrTooShortPassword = fmt.Errorf("a senha precisa conter ao menos %d caracters", minPasswordLength)

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
}

func NewUser(name, email, password string) (User, error) {
	pwd, err := generatePassword(password)
	if err != nil {
		return User{}, err
	}

	u := User{
		ID:           uuid.New(),
		Name:         strings.TrimSpace(name),
		Email:        strings.TrimSpace(strings.ToLower(email)),
		PasswordHash: pwd,
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
