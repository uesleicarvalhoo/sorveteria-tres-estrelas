//go:build unit || all

package auth_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/user"
	userMock "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/user/mocks"
)

const (
	secretKey = "my-super-secret-key"
	issuer    = "test"
)

func TestLogin(t *testing.T) {
	t.Parallel()

	password := "my-secret-password"
	storedUser, _ := user.NewUser("User LastName", "user.lastname@email.com", password)

	t.Run("test valid", func(t *testing.T) {
		t.Parallel()

		// Arrange
		provider := auth.NewStaticProvider(secretKey, issuer)

		repo := userMock.NewRepository(t)
		repo.On("GetByEmail", mock.Anything, storedUser.Email).Return(storedUser, nil).Once()

		startAt := time.Now()

		userUc := user.NewService(repo)

		sut := auth.NewService(userUc, provider)

		// Action
		token, err := sut.Login(context.Background(), auth.LoginPayload{Email: storedUser.Email, Password: password})

		// Assert
		assert.NoError(t, err)

		assert.Equal(t, "bearer", token.GrantType)
		assert.Greater(t, token.ExpiresAt, startAt.Unix())
		assert.NotEmpty(t, token.Token)

		u, err := sut.Authorize(context.Background(), token.Token)
		assert.Equal(t, storedUser.ID, u.ID)
		assert.Equal(t, storedUser.Name, u.Name)
		assert.Equal(t, storedUser.Email, u.Email)
		assert.NoError(t, err)
	})

	t.Run("test errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			about           string
			email           string
			password        string
			repositoryUser  user.User
			repositoryError error
			expectedError   string
		}{
			{
				about:           "when user is not found",
				email:           "inexisting@email.com",
				repositoryError: errors.New("user not found"),
				expectedError:   "user not found",
			},
			{
				about:          "when password is invalid",
				email:          storedUser.Email,
				password:       "wrongPassword",
				repositoryUser: storedUser,
				expectedError:  auth.ErrNotAuthorized.Error(),
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()

				// Arrange
				provider := auth.NewStaticProvider(secretKey, issuer)

				repo := userMock.NewRepository(t)
				repo.On("GetByEmail", mock.Anything, tc.email).Return(tc.repositoryUser, tc.repositoryError).Once()

				sut := auth.NewService(user.NewService(repo), provider)

				// Action
				token, err := sut.Login(context.Background(), auth.LoginPayload{tc.email, tc.password})

				// Assert
				assert.EqualError(t, err, tc.expectedError)
				assert.Equal(t, auth.JwtToken{}, token)
			})
		}
	})
}

func TestRefreshToken(t *testing.T) {
	t.Parallel()

	t.Run("when token is valid", func(t *testing.T) {
		t.Parallel()

		// Arrange
		storedUser := user.User{ID: uuid.New()}

		token, err := auth.GenerateJwtToken(context.Background(), storedUser, time.Now().Add(time.Hour), issuer, secretKey)
		assert.NoError(t, err)

		sut := auth.NewService(user.NewService(userMock.NewRepository(t)), auth.NewStaticProvider(secretKey, issuer))

		// Action
		time.Sleep(time.Second) // Wait 1 second for change token
		newToken, err := sut.RefreshToken(context.Background(), auth.RefreshTokenPayload{token})

		// Assert
		assert.NoError(t, err)
		assert.NotEqual(t, token, newToken.Token)
		assert.Greater(t, newToken.ExpiresAt, time.Now().Unix())
	})
}

func TestAuthorize(t *testing.T) {
	t.Parallel()

	t.Run("when token is valid", func(t *testing.T) {
		t.Parallel()

		// Arrange
		storedUser, err := user.NewUser("User Name", "user@email.com.br", "secret123")
		assert.NoError(t, err)

		token, err := auth.GenerateJwtToken(context.Background(), storedUser, time.Now().Add(time.Hour), issuer, secretKey)
		assert.NoError(t, err)

		mockUserSvc := userMock.NewUseCase(t)

		sut := auth.NewService(mockUserSvc, auth.NewStaticProvider(secretKey, issuer))

		// Action
		user, err := sut.Authorize(context.Background(), token)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, storedUser.ID, user.ID)
		assert.Equal(t, storedUser.Name, user.Name)
		assert.Equal(t, storedUser.Email, user.Email)
	})
}
