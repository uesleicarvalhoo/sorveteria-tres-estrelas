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
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	cacheMocks "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cache/mocks"
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

		accessKey := fmt.Sprintf("access-token-%s", storedUser.ID.String())
		refreshKey := fmt.Sprintf("refresh-token-%s", storedUser.ID.String())

		mockCache := cacheMocks.NewCache(t)
		mockCache.On("Set", mock.Anything, accessKey, mock.Anything, auth.AccessTokenDuration).Return(nil).Once()
		mockCache.On("Set", mock.Anything, refreshKey, mock.Anything, auth.RefreshTokenDuration).Return(nil).Once()

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

				accessKey := fmt.Sprintf("access-token-%s", storedUser.ID.String())
				refreshKey := fmt.Sprintf("refresh-token-%s", storedUser.ID.String())

				mockCache := cacheMocks.NewCache(t)
				mockCache.On("Set", mock.Anything, accessKey, mock.Anything, auth.AccessTokenDuration).
					Return(tc.cacheError).Maybe()
				mockCache.On("Set", mock.Anything, refreshKey, mock.Anything, auth.RefreshTokenDuration).
					Return(tc.cacheError).Maybe()

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
		userID := uuid.New()

		token, err := auth.GenerateJwtToken(secretKey, userID, time.Now().Add(time.Hour))
		assert.NoError(t, err)

		accessKey := fmt.Sprintf("access-token-%s", userID.String())
		refreshKey := fmt.Sprintf("refresh-token-%s", userID.String())

		mockCache := cacheMocks.NewCache(t)
		mockCache.On("Set", mock.Anything, accessKey, mock.Anything, auth.AccessTokenDuration).Return(nil).Once()
		mockCache.On("Set", mock.Anything, refreshKey, mock.Anything, auth.RefreshTokenDuration).Return(nil).Once()
		mockCache.On("Get", mock.Anything, refreshKey).Return(token, nil).Once()

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
		userID := uuid.New()

		token, err := auth.GenerateJwtToken(secretKey, userID, time.Now().Add(time.Hour))
		assert.NoError(t, err)

		refreshTokenKey := fmt.Sprintf("refresh-token-%s", userID.String())

		mockCache := cacheMocks.NewCache(t)
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
		userID := uuid.New()

		token, err := auth.GenerateJwtToken(secretKey, userID, time.Now().Add(time.Hour))
		assert.NoError(t, err)

		refreshTokenKey := fmt.Sprintf("refresh-token-%s", userID.String())
		mockError := errors.New("cache error")

		mockCache := cacheMocks.NewCache(t)
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
		storedUser, err := entity.NewUser("User Name", "user@email.com.br", "secret123", entity.ReadWritePopsicles, entity.ReadWriteSales)

		token, err := auth.GenerateJwtToken(secretKey, storedUser.ID, time.Now().Add(time.Hour))
		assert.NoError(t, err)

		accessTokenKey := fmt.Sprintf("access-token-%s", storedUser.ID.String())

		mockCache := cacheMocks.NewCache(t)
		mockCache.On("Get", mock.Anything, accessTokenKey).Return(token, nil).Once()

		mockUserSvc := userMocks.NewUseCase(t)
		mockUserSvc.On("Get", mock.Anything, storedUser.ID).Return(storedUser, nil)

		sut := auth.NewService(secretKey, mockUserSvc, mockCache)

		// Action
		sub, err := sut.Authorize(context.Background(), token, "popsicles", "write")

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, storedUser.ID, sub)
	})

	t.Run("check errors", func(t *testing.T) {
		t.Parallel()

		storedUser, err := entity.NewUser("User Name", "user@email.com.br", "secret123", entity.ReadWritePopsicles, entity.ReadSales)

		validToken, err := auth.GenerateJwtToken(secretKey, storedUser.ID, time.Now().Add(time.Hour))
		assert.NoError(t, err)

		tests := []struct {
			about           string
			mockUserError   error
			mockCacheError  error
			mockCacheReturn string
			expectedError   string
			userPermissions []entity.Permission
			domain          string
			action          string
		}{
			{
				about:           "when token is valid but not match with cached token",
				mockCacheReturn: "wrong-token",
				expectedError:   auth.ErrTokenNotFound.Error(),
			},
			{
				about:           "when token is valid but not match with cached token",
				mockCacheReturn: "",
				mockCacheError:  errors.New("cache error"),
				expectedError:   "cache error",
			},
			{
				about:           "when token is valid but not match with cached token",
				mockCacheReturn: "",
				mockCacheError:  errors.New("cache error"),
				expectedError:   "cache error",
			},
			{
				about:           "when user don't have permission",
				mockCacheReturn: validToken,
				mockCacheError:  nil,
				domain:          "sales",
				action:          "write",
				expectedError:   auth.ErrNotPermited.Error(),
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()

				// Arrange
				accessTokenKey := fmt.Sprintf("access-token-%s", storedUser.ID.String())

				mockCache := cacheMocks.NewCache(t)
				mockCache.On("Get", mock.Anything, accessTokenKey).Return(tc.mockCacheReturn, tc.mockCacheError).Once()

				mockUserSvc := userMocks.NewUseCase(t)
				mockUserSvc.On("Get", mock.Anything, storedUser.ID).Return(storedUser, err).Maybe()

				sut := auth.NewService(secretKey, mockUserSvc, mockCache)

				// Action
				sub, err := sut.Authorize(context.Background(), validToken, tc.domain, tc.action)

				// Assert
				assert.Equal(t, uuid.Nil, sub)
				assert.EqualError(t, err, tc.expectedError)
			})
		}
	})
}
