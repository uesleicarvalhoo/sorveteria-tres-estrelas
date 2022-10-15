//go:build unit || all

package auth_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth/mocks"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/user"
	userMocks "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/user/mocks"
)

const secretKey = "my-super-secret-key"

func TestLogin(t *testing.T) {
	t.Parallel()

	password := "my-secret-password"
	storedUser, _ := entity.NewUser("User LastName", "user.lastname@email.com", password)

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
		token, err := sut.Login(context.Background(), storedUser.Email, password)

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
			email           string
			password        string
			repositoryUser  entity.User
			repositoryError error
			cacheError      error
			expectedError   string
		}{
			{
				about:           "when user is not found",
				email:           "inexisting@email.com",
				repositoryError: errors.New("user not found"),
				expectedError:   "user not found",
			},
			{
				about:          "when cache return an error",
				email:          storedUser.Email,
				password:       password,
				repositoryUser: storedUser,
				cacheError:     errors.New("couldn't set value into cache"),
				expectedError:  "couldn't set value into cache",
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
				repo := userMocks.NewRepository(t)
				repo.On("GetByEmail", mock.Anything, tc.email).Return(tc.repositoryUser, tc.repositoryError).Once()

				accessTokenKey := fmt.Sprintf("access-token-%s", storedUser.ID.String())
				refreshTokenKey := fmt.Sprintf("refresh-token-%s", storedUser.ID.String())

				mockCache := mocks.NewCache(t)
				mockCache.On("Set", mock.Anything, accessTokenKey, mock.Anything).Return(tc.cacheError).Maybe()
				mockCache.On("Set", mock.Anything, refreshTokenKey, mock.Anything).Return(tc.cacheError).Maybe()

				sut := auth.NewService(secretKey, user.NewService(repo), mockCache)

				// Action
				token, err := sut.Login(context.Background(), tc.email, tc.password)

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
		t.Parallel()

		// Arrange
		userID := entity.NewID()

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
		t.Parallel()

		// Arrange
		userID := entity.NewID()

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
		t.Parallel()

		// Arrange
		userID := entity.NewID()

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
		t.Parallel()

		// Arrange
		userID := entity.NewID()

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
		t.Parallel()

		// Arrange
		userID := entity.NewID()

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
		assert.Equal(t, entity.ID{}, sub)
	})

	t.Run("when cache return an error", func(t *testing.T) {
		t.Parallel()

		// Arrange

		userID := entity.NewID()

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
		assert.Equal(t, entity.ID{}, sub)
	})
}
