//go:build unit || all

package user_test

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/user"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	t.Run("should convert email to lower case", func(t *testing.T) {
		t.Parallel()
		// Arrange
		name := "Ueslei Carvalho"
		email := "UESLEICDOLIVEIRA@gmail.com"
		passwd := "secret123"

		// Action
		u, err := user.NewUser(name, email, passwd)

		// Assert
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, u.ID)
		assert.Equal(t, strings.ToLower(email), u.Email)
		assert.NotEqual(t, passwd, u.PasswordHash)
		assert.True(t, u.CheckPassword(passwd))
		assert.False(t, u.CheckPassword("wrong-password"))
	})

	t.Run("errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			about       string
			name        string
			email       string
			passwd      string
			expectedErr string
		}{
			{
				about:       "when name is empty",
				name:        "",
				email:       "ueslei.carvalho@email.com",
				passwd:      "mySecretPassword!",
				expectedErr: "nome: campo obrigatório",
			},
			{
				about:       "when email is empty",
				name:        "Ueslei Carvalho",
				email:       "",
				passwd:      "mySecretPassword!",
				expectedErr: "email: campo invalido",
			},
			{
				about:       "when email is invalid",
				name:        "Ueslei Carvalho",
				email:       "wrongemail",
				passwd:      "mySecretPassword!",
				expectedErr: "email: campo invalido",
			},

			{
				about:       "when password has lower then 5 caracters",
				name:        "User Lastname",
				passwd:      "1234",
				expectedErr: "a senha precisa conter ao menos 5 caracters",
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()

				// Action
				u, err := user.NewUser(tc.name, tc.email, tc.passwd)

				// Assert
				assert.Equal(t, user.User{}, u)
				assert.EqualError(t, err, tc.expectedErr)
			})
		}
	})
}

func TestUserValidate(t *testing.T) {
	t.Parallel()

	t.Run("check a valid user", func(t *testing.T) {
		t.Parallel()

		// Arrange
		u := user.User{
			ID:           uuid.New(),
			Name:         "Ueslei Carvalho",
			Email:        "uesleicdoliveira@gmail.com",
			PasswordHash: "mySecretPassword!",
		}

		// Action
		err := u.Validate()

		// Assert
		assert.NoError(t, err)
	})

	tests := []struct {
		about         string
		user          user.User
		expectedError string
	}{
		{
			about: "when name is empty",
			user: user.User{
				ID:           uuid.New(),
				Email:        "uesleicdoliveira@gmail.com",
				PasswordHash: "mySecretPassword!",
			},
			expectedError: "nome: campo obrigatório",
		},
		{
			about: "when email is empty",
			user: user.User{
				ID:           uuid.New(),
				Name:         "Ueslei Carvalho",
				PasswordHash: "mySecretPassword!",
			},
			expectedError: "email: campo invalido",
		},
		{
			about: "when email is invalid",
			user: user.User{
				ID:           uuid.New(),
				Name:         "Ueslei Carvalho",
				Email:        "wrong!",
				PasswordHash: "mySecretPassword!",
			},
			expectedError: "email: campo invalido",
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			err := tc.user.Validate()

			assert.EqualError(t, err, tc.expectedError)
		})
	}
}
