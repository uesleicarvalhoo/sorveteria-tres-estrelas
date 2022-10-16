//go:build unit || all

package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/user"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/user/mocks"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	t.Run("when all fields are ok", func(t *testing.T) {
		t.Parallel()

		// Arrange
		repo := mocks.NewRepository(t)
		repo.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

		sut := user.NewService(repo)

		name := "Ueslei Carvalho"
		email := "ueslei.carvalho@email.com"
		password := "secret123"
		permissions := []entity.Permission{entity.ReadWritePopsicles, entity.ReadWriteSales}

		// Action
		u, err := sut.Create(context.Background(), name, email, password, permissions...)

		// Assert
		assert.NoError(t, err)

		assert.NotEqual(t, uuid.Nil, u.ID)
		assert.Equal(t, email, u.Email)
		assert.NotEqual(t, password, u.PasswordHash)
		assert.True(t, u.CheckPassword(password))
		assert.Equal(t, permissions, u.Permissions)
	})

	tests := []struct {
		about         string
		name          string
		email         string
		password      string
		permissions   []entity.Permission
		repoError     error
		expectedError string
	}{
		{
			about:         "when repository return an error",
			name:          "Ueslei Carvalho",
			email:         "ueslei.carvalho@email.com",
			password:      "mySecretPassword!",
			repoError:     errors.New("failed to create a new user"),
			expectedError: "failed to create a new user",
		},
		{
			about:         "when name is empty",
			name:          "",
			email:         "ueslei.carvalho@email.com",
			password:      "mySecretPassword!",
			expectedError: "Name é obrigatorio",
		},
		{
			about:         "when email is empty",
			name:          "Ueslei Carvalho",
			email:         "",
			password:      "mySecretPassword!",
			expectedError: "'' não é um email valido",
		},
		{
			about:         "when email is invalid",
			name:          "Ueslei Carvalho",
			email:         "wrongemail",
			password:      "mySecretPassword!",
			expectedError: "'wrongemail' não é um email valido",
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			repo := mocks.NewRepository(t)
			repo.On("Create", mock.Anything, mock.Anything).Return(tc.repoError).Maybe()

			sut := user.NewService(repo)
			// Action
			u, err := sut.Create(context.Background(), tc.name, tc.email, tc.password, tc.permissions...)

			// Assert
			assert.Equal(t, entity.User{}, u)
			assert.EqualError(t, err, tc.expectedError)
		})
	}
}

func TestGet(t *testing.T) {
	t.Parallel()

	storedUser, _ := entity.NewUser("Name LastName", "user@email.com.br", "fakehash:123")

	tests := []struct {
		about        string
		id           uuid.UUID
		expectedUser entity.User
		mockError    error
		expectedErr  error
	}{
		{
			about:       "when repository return an error should return error and an empty user",
			id:          storedUser.ID,
			mockError:   errors.New("record not found"),
			expectedErr: errors.New("record not found"),
		},
		{
			about:        "when repository return a user",
			id:           storedUser.ID,
			expectedUser: storedUser,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			repo := mocks.NewRepository(t)
			repo.On("Get", mock.Anything, storedUser.ID).Return(storedUser, tc.mockError).Once()

			sut := user.NewService(repo)

			// Action
			found, err := sut.Get(context.Background(), tc.id)

			// Assert
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedUser, found)
		})

	}
}

func TestGetByEmail(t *testing.T) {
	t.Parallel()

	storedUser, _ := entity.NewUser("Name LastName", "user@email.com.br", "fakehash:123")

	tests := []struct {
		about        string
		email        string
		expectedUser entity.User
		mockError    error
		expectedErr  error
	}{
		{
			about:       "when repository return an error should return error and an empty user",
			email:       storedUser.Email,
			mockError:   errors.New("record not found"),
			expectedErr: errors.New("record not found"),
		},
		{
			about:        "when repository return a user",
			email:        storedUser.Email,
			expectedUser: storedUser,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			repo := mocks.NewRepository(t)
			repo.On("GetByEmail", mock.Anything, storedUser.Email).Return(storedUser, tc.mockError).Once()

			sut := user.NewService(repo)

			// Action
			found, err := sut.GetByEmail(context.Background(), tc.email)

			// Assert
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedUser, found)
		})

	}
}
