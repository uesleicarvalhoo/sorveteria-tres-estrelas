// go:build unit || all

package users_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	t.Run("when permissions are empty", func(t *testing.T) {
		t.Parallel()
		// Arrange
		name := "Ueslei Carvalho"
		email := "uesleicdoliveira@gmail.com"
		passwd := "secret123"

		// Action
		user, err := users.NewUser(name, email, passwd)

		// Assert
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, user.ID)
		assert.Equal(t, email, user.Email)
		assert.NotEqual(t, passwd, user.PasswordHash)
		assert.True(t, user.CheckPassword(passwd))
		assert.False(t, user.CheckPassword("wrong-password"))
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
				u, err := users.NewUser(tc.name, tc.email, tc.passwd)

				// Assert
				assert.Equal(t, users.User{}, u)
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
		u := users.User{
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
		user          users.User
		expectedError string
	}{
		{
			about: "when name is empty",
			user: users.User{
				ID:           uuid.New(),
				Email:        "uesleicdoliveira@gmail.com",
				PasswordHash: "mySecretPassword!",
			},
			expectedError: "nome: campo obrigatório",
		},
		{
			about: "when email is empty",
			user: users.User{
				ID:           uuid.New(),
				Name:         "Ueslei Carvalho",
				PasswordHash: "mySecretPassword!",
			},
			expectedError: "email: campo invalido",
		},
		{
			about: "when email is invalid",
			user: users.User{
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
