//go:build unit || all

package auth_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

func TestJWT(t *testing.T) {
	t.Parallel()

	userID := entity.NewID()
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
		assert.Equal(t, entity.ID{}, sub)
	})
}
