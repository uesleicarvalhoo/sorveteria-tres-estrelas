//go:build unit || all

package auth_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/auth/mocks"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user"
	userMocks "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user/mocks"
)

const secretKey = "my-super-secret-key"

func TestCreateUser(t *testing.T) {
	t.Parallel()

	t.Run("password should be encrypted", func(t *testing.T) {
		t.Parallel()

		// Arrange
		repo := userMocks.NewRepository(t)
		repo.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

		cacheMock := mocks.NewCache(t)

		sut := auth.NewService(secretKey, user.NewService(repo), cacheMock)
		payload := dto.CreateUserPayload{
			Name:     "Fake Lastname",
			Email:    "fake.lastname@email.com",
			Password: "secret123",
		}

		// Action
		u, err := sut.CreateUser(context.Background(), payload)

		// Assert
		assert.NoError(t, err)

		assert.True(t, auth.CheckPasswordHash(payload.Password, u.PasswordHash))
		assert.False(t, auth.CheckPasswordHash("wrongPassword", u.PasswordHash))
	})

	t.Run("should create with default permissions", func(t *testing.T) {
		t.Parallel()

		// Arrange
		repo := userMocks.NewRepository(t)
		repo.On("Create", mock.Anything, mock.Anything).Return(nil).Once()
		cacheMock := mocks.NewCache(t)

		sut := auth.NewService(secretKey, user.NewService(repo), cacheMock)
		payload := dto.CreateUserPayload{
			Name:     "Fake Lastname",
			Email:    "fake.lastname@email.com",
			Password: "secret123",
		}

		// Action
		u, err := sut.CreateUser(context.Background(), payload)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, auth.GetDefaultUserPermissions(), u.Permissions)
	})

	t.Run("when user.UseCase return an error", func(t *testing.T) {
		t.Parallel()

		// Arrange
		mockErr := errors.New("failed to create a new user")

		repo := userMocks.NewRepository(t)
		repo.On("Create", mock.Anything, mock.Anything).Return(mockErr).Once()
		cacheMock := mocks.NewCache(t)

		sut := auth.NewService(secretKey, user.NewService(repo), cacheMock)
		payload := dto.CreateUserPayload{
			Name:     "Fake Lastname",
			Email:    "fake.lastname@email.com",
			Password: "secret123",
		}

		// Action
		u, err := sut.CreateUser(context.Background(), payload)

		// Assert
		assert.Equal(t, user.User{}, u)
		assert.EqualError(t, err, mockErr.Error())
	})
}

func TestLogin(t *testing.T) {
	t.Parallel()

	payload := dto.LoginPayload{
		Email:    "user.lastname@email.com.br",
		Password: "my-secret-password",
	}

	passwdHash, err := auth.GeneratePasswordHash(payload.Password)
	assert.NoError(t, err)

	storedUser := user.User{
		ID:           uuid.New(),
		Name:         "User Lastname",
		Email:        payload.Email,
		PasswordHash: passwdHash,
		Permissions:  auth.GetDefaultUserPermissions(),
	}

	t.Run("test valid", func(t *testing.T) {
		t.Parallel()

		// Arrange
		repo := userMocks.NewRepository(t)
		repo.On("GetByEmail", mock.Anything, storedUser.Email).Return(storedUser, nil).Once()

		accessTokenKey := fmt.Sprintf("access-token-%s", storedUser.ID.String())
		refreshTokenKey := fmt.Sprintf("refresh-token-%s", storedUser.ID.String())

		mockCache := mocks.NewCache(t)
		mockCache.On("Set", mock.Anything, accessTokenKey, mock.Anything).Return(nil).Once()
		mockCache.On("Set", mock.Anything, refreshTokenKey, mock.Anything).Return(nil).Once()

		startAt := time.Now()

		userUc := user.NewService(repo)

		sut := auth.NewService(secretKey, userUc, mockCache)

		// Action
		token, err := sut.Login(context.Background(), payload)

		// Assert
		assert.NoError(t, err)

		assert.Equal(t, "bearer", token.GrantType)
		assert.Greater(t, token.ExpiresAt, startAt.Unix())

		assert.NotEmpty(t, token.AcessToken)
		assert.NotEmpty(t, token.RefreshToken)
		assert.NotEqual(t, token.AcessToken, token.RefreshToken)
	})

	t.Run("test errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			about           string
			payload         dto.LoginPayload
			repositoryUser  user.User
			repositoryError error
			cacheError      error
			expectedError   string
		}{
			{
				about:           "when user is not found",
				payload:         payload,
				repositoryError: errors.New("user not found"),
				expectedError:   "user not found",
			},
			{
				about:          "when cache return an error",
				payload:        payload,
				repositoryUser: storedUser,
				cacheError:     errors.New("couldn't set value into cache"),
				expectedError:  "couldn't set value into cache",
			},
			{
				about:          "when password is invalid",
				payload:        dto.LoginPayload{Email: payload.Email, Password: "wrongPassword"},
				repositoryUser: storedUser,
				expectedError:  auth.ErrNotAuthorized.Error(),
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()

				// Arrange
				repo := userMocks.NewRepository(t)
				repo.On("GetByEmail", mock.Anything, storedUser.Email).Return(tc.repositoryUser, tc.repositoryError).Once()

				accessTokenKey := fmt.Sprintf("access-token-%s", storedUser.ID.String())
				refreshTokenKey := fmt.Sprintf("refresh-token-%s", storedUser.ID.String())

				mockCache := mocks.NewCache(t)
				mockCache.On("Set", mock.Anything, accessTokenKey, mock.Anything).Return(tc.cacheError).Maybe()
				mockCache.On("Set", mock.Anything, refreshTokenKey, mock.Anything).Return(tc.cacheError).Maybe()

				sut := auth.NewService(secretKey, user.NewService(repo), mockCache)

				// Action
				token, err := sut.Login(context.Background(), tc.payload)

				// Assert
				assert.EqualError(t, err, tc.expectedError)
				assert.Equal(t, auth.JwtToken{}, token)
			})
		}
	})
}

func TestRefreshToken(t *testing.T) {
	t.Parallel()

	t.Run("when token is valid and cache not return an error", func(t *testing.T) {
		// Arrange
		userID := uuid.New()

		token, err := auth.GenerateJwtToken(secretKey, userID, time.Now().Add(time.Hour))
		assert.NoError(t, err)

		accessTokenKey := fmt.Sprintf("access-token-%s", userID.String())
		refreshTokenKey := fmt.Sprintf("refresh-token-%s", userID.String())

		mockCache := mocks.NewCache(t)
		mockCache.On("Set", mock.Anything, accessTokenKey, mock.Anything).Return(nil).Once()
		mockCache.On("Set", mock.Anything, refreshTokenKey, mock.Anything).Return(nil).Once()
		mockCache.On("Get", mock.Anything, refreshTokenKey).Return(token, nil).Once()

		sut := auth.NewService(secretKey, user.NewService(userMocks.NewRepository(t)), mockCache)

		// Action
		time.Sleep(time.Second) // Wait 1 second for change token
		newToken, err := sut.RefreshToken(context.Background(), token)

		// Assert
		assert.NoError(t, err)
		assert.NotEqual(t, token, newToken.AcessToken)
		assert.NotEqual(t, token, newToken.RefreshToken)
		assert.Greater(t, newToken.ExpiresAt, time.Now().Unix())
	})

	t.Run("when token is valid but not match with cached token", func(t *testing.T) {
		// Arrange
		userID := uuid.New()

		token, err := auth.GenerateJwtToken(secretKey, userID, time.Now().Add(time.Hour))
		assert.NoError(t, err)

		refreshTokenKey := fmt.Sprintf("refresh-token-%s", userID.String())

		mockCache := mocks.NewCache(t)
		mockCache.On("Get", mock.Anything, refreshTokenKey).Return("wrong-token", nil).Once()

		sut := auth.NewService(secretKey, user.NewService(userMocks.NewRepository(t)), mockCache)

		// Action
		time.Sleep(time.Second) // Wait 1 second for change token
		newToken, err := sut.RefreshToken(context.Background(), token)

		// Assert
		assert.EqualError(t, err, auth.ErrTokenNotFound.Error())
		assert.Equal(t, auth.JwtToken{}, newToken)
	})

	t.Run("when cache return an error", func(t *testing.T) {
		// Arrange
		userID := uuid.New()

		token, err := auth.GenerateJwtToken(secretKey, userID, time.Now().Add(time.Hour))
		assert.NoError(t, err)

		refreshTokenKey := fmt.Sprintf("refresh-token-%s", userID.String())
		mockError := errors.New("cache error")

		mockCache := mocks.NewCache(t)
		mockCache.On("Get", mock.Anything, refreshTokenKey).Return("", mockError).Once()

		sut := auth.NewService(secretKey, user.NewService(userMocks.NewRepository(t)), mockCache)

		// Action
		time.Sleep(time.Second) // Wait 1 second for change token
		newToken, err := sut.RefreshToken(context.Background(), token)

		// Assert
		assert.EqualError(t, err, mockError.Error())
		assert.Equal(t, auth.JwtToken{}, newToken)
	})
}

func TestAuthorize(t *testing.T) {
	t.Parallel()

	t.Run("when token is valid and cache not return an error", func(t *testing.T) {
		// Arrange
		userID := uuid.New()

		token, err := auth.GenerateJwtToken(secretKey, userID, time.Now().Add(time.Hour))
		assert.NoError(t, err)

		accessTokenKey := fmt.Sprintf("access-token-%s", userID.String())

		mockCache := mocks.NewCache(t)
		mockCache.On("Get", mock.Anything, accessTokenKey).Return(token, nil).Once()

		sut := auth.NewService(secretKey, user.NewService(userMocks.NewRepository(t)), mockCache)

		// Action
		sub, err := sut.Authorize(context.Background(), token)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, userID, sub)
	})

	t.Run("when token is valid but not match with cached token", func(t *testing.T) {
		// Arrange
		userID := uuid.New()

		token, err := auth.GenerateJwtToken(secretKey, userID, time.Now().Add(time.Hour))
		assert.NoError(t, err)

		accessTokenKey := fmt.Sprintf("access-token-%s", userID.String())

		mockCache := mocks.NewCache(t)
		mockCache.On("Get", mock.Anything, accessTokenKey).Return("wrong-token", nil).Once()

		sut := auth.NewService(secretKey, user.NewService(userMocks.NewRepository(t)), mockCache)

		// Action
		sub, err := sut.Authorize(context.Background(), token)

		// Assert
		assert.EqualError(t, err, auth.ErrTokenNotFound.Error())
		assert.Equal(t, uuid.Nil, sub)
	})

	t.Run("when cache return an error", func(t *testing.T) {
		// Arrange
		userID := uuid.New()

		token, err := auth.GenerateJwtToken(secretKey, userID, time.Now().Add(time.Hour))
		assert.NoError(t, err)

		accessTokenKey := fmt.Sprintf("access-token-%s", userID.String())
		mockError := errors.New("cache error")

		mockCache := mocks.NewCache(t)
		mockCache.On("Get", mock.Anything, accessTokenKey).Return("", mockError).Once()

		sut := auth.NewService(secretKey, user.NewService(userMocks.NewRepository(t)), mockCache)

		// Action
		sub, err := sut.Authorize(context.Background(), token)

		// Assert
		assert.EqualError(t, err, mockError.Error())
		assert.Equal(t, uuid.Nil, sub)
	})
}
