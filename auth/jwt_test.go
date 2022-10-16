//go:build unit || all

package auth_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
)

func TestJWT(t *testing.T) {
	t.Parallel()

	userID := uuid.New()
	secret := "my-secret-key"

	t.Run("when token is valid", func(t *testing.T) {
		t.Parallel()
		// Arrange
		exp := time.Now().Add(time.Minute)

		token, err := auth.GenerateJwtToken(secret, userID, exp)
		assert.NoError(t, err)

		// Action
		sub, err := auth.ValidateJwtToken(token, secret)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, userID, sub)
	})

	t.Run("when token is expired", func(t *testing.T) {
		t.Parallel()

		// Arrange
		exp := time.Now().Add(-time.Minute)

		token, err := auth.GenerateJwtToken(secret, userID, exp)
		assert.NoError(t, err)

		// Action
		sub, err := auth.ValidateJwtToken(token, secret)

		// Assert
		assert.EqualError(t, err, "Token is expired")
		assert.Equal(t, uuid.Nil, sub)
	})
}
