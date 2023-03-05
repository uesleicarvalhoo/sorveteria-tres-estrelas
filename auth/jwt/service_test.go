package jwt_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth/jwt"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user"
)

func TestJWT(t *testing.T) {
	t.Parallel()

	storedUser, err := user.NewUser("Username", "user@email.com", "imsecret")
	assert.NoError(t, err)

	secret := "my-secret-key"

	t.Run("when token is valid", func(t *testing.T) {
		t.Parallel()

		// Arrange
		sut := jwt.NewService(secret)
		exp := time.Now().Add(time.Minute)

		token, err := sut.Generate(context.Background(), storedUser, exp)
		assert.NoError(t, err)

		// Action
		tokenizedUser, err := sut.Validate(context.Background(), token)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, storedUser, tokenizedUser)
	})

	t.Run("when token is expired", func(t *testing.T) {
		t.Parallel()

		// Arrange
		sut := jwt.NewService(secret)
		exp := time.Now().Add(-time.Minute)

		token, err := sut.Generate(context.Background(), storedUser, exp)
		assert.NoError(t, err)

		// Action
		sub, err := sut.Validate(context.Background(), token)

		// Assert
		assert.EqualError(t, err, "Token is expired")
		assert.Equal(t, user.User{}, sub)
	})
}
