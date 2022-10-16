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

	t.Run("when permissions are empty", func(t *testing.T) {
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
		assert.Equal(t, entity.DefaultPermissions(), user.Permissions)
	})

	t.Run("when permissions has permissions", func(t *testing.T) {
		t.Parallel()
		// Arrange
		name := "Ueslei Carvalho"
		email := "uesleicdoliveira@gmail.com"
		passwd := "secret123"
		permissions := []entity.Permission{entity.ReadWritePopsicles, entity.ReadWriteSales, entity.ReadWriteUsers}

		// Action
		user, err := entity.NewUser(name, email, passwd, permissions...)

		// Assert
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, user.ID)
		assert.Equal(t, email, user.Email)
		assert.NotEqual(t, passwd, user.PasswordHash)
		assert.True(t, user.CheckPassword(passwd))
		assert.False(t, user.CheckPassword("wrong-password"))
		assert.Equal(t, permissions, user.Permissions)
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

func TestUserCheckPermissions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		about          string
		permissions    []entity.Permission
		domain         string
		action         string
		expectedResult bool
	}{
		{
			about:          "when user has permissions to make action on domain",
			permissions:    []entity.Permission{entity.ReadWritePopsicles},
			domain:         "popsicles",
			action:         "write",
			expectedResult: true,
		},
		{
			about:          "when user not have permission to make action on domain",
			permissions:    []entity.Permission{entity.ReadPopsicles},
			domain:         "popsicles",
			action:         "write",
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			user, err := entity.NewUser("user name", "user@email.com", "secret-123", tc.permissions...)
			assert.NoError(t, err)

			// Action
			ok := user.AuthorizeDomainAction(tc.domain, tc.action)

			// Assert
			assert.Equal(t, tc.expectedResult, ok)
		})
	}
}
