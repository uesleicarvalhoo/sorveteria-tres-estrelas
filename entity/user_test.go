// go:build unit || all

package entity_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	t.Run("when all fields are ok", func(t *testing.T) {
		t.Parallel()
		// Arrange
		name := "Ueslei Carvalho"
		email := "uesleicdoliveira@gmail.com"
		passwd := "secret123"

		// Action
		user, err := entity.NewUser(name, email, passwd)

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
			permissions []entity.Permission
			expectedErr string
		}{
			{
				about:       "when name is empty",
				name:        "",
				email:       "ueslei.carvalho@email.com",
				passwd:      "mySecretPassword!",
				expectedErr: "Name é obrigatorio",
			},
			{
				about:       "when email is empty",
				name:        "Ueslei Carvalho",
				email:       "",
				passwd:      "mySecretPassword!",
				expectedErr: "'' não é um email valido",
			},
			{
				about:       "when email is invalid",
				name:        "Ueslei Carvalho",
				email:       "wrongemail",
				passwd:      "mySecretPassword!",
				expectedErr: "'wrongemail' não é um email valido",
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
				u, err := entity.NewUser(tc.name, tc.email, tc.passwd, tc.permissions...)

				// Assert
				assert.Equal(t, entity.User{}, u)
				assert.EqualError(t, err, tc.expectedErr)
			})
		}
	})
}
