package password

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const MinPasswordLength = 5

var ErrTooShortPassword = fmt.Errorf("a senha precisa conter %d ou mais caracters", MinPasswordLength)

type BCrypt struct{}

func NewBCrypt() *BCrypt {
	return &BCrypt{}
}

func (BCrypt) GenerateHash(password string) (string, error) {
	if len(password) < MinPasswordLength {
		return "", ErrTooShortPassword
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

func (BCrypt) CheckHash(plain, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))

	return err == nil
}
