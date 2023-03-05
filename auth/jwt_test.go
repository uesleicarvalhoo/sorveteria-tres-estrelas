package auth_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user"
)

func TestJWT(t *testing.T) {
	t.Parallel()

	storedUser, err := user.NewUser("Username", "user@email.com", "imsecret")
	assert.NoError(t, err)

	issuer := "issuer"
	secret := "my-secret-key"

	t.Run("when token is valid", func(t *testing.T) {
		t.Parallel()

		// Arrange
		exp := time.Now().Add(time.Minute)

		token, err := auth.GenerateJwtToken(context.Background(), storedUser, exp, issuer, secret)
		assert.NoError(t, err)

		// Action
		tokenizedUser, err := auth.ValidateJwtToken(context.Background(), token, secret)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, storedUser.ID, tokenizedUser.ID)
		assert.Equal(t, storedUser.Name, tokenizedUser.Name)
		assert.Equal(t, storedUser.Email, tokenizedUser.Email)
		assert.Equal(t, tokenizedUser.PasswordHash, "")
	})

	t.Run("when token is expired", func(t *testing.T) {
		t.Parallel()

		// Arrange
		exp := time.Now().Add(-time.Minute)

		token, err := auth.GenerateJwtToken(context.Background(), storedUser, exp, issuer, secret)
		assert.NoError(t, err)

		// Action
		sub, err := auth.ValidateJwtToken(context.Background(), token, secret)

		// Assert
		assert.EqualError(t, err, "Token is expired")
		assert.Equal(t, user.User{}, sub)
	})
}
