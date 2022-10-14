//go:build unit || all

package auth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/auth"
)

func TestGeneratePassword(t *testing.T) {
	t.Parallel()

	t.Run("invalid password", func(t *testing.T) {
		t.Parallel()

		// Arrange
		password := ""

		// Action
		hash, err := auth.GeneratePasswordHash(password)

		// Assert
		assert.Equal(t, "", hash)
		assert.EqualError(t, err, auth.ErrTooShortPassword.Error())
	})

	t.Run("valid password", func(t *testing.T) {
		t.Parallel()

		// Arrange
		password := "strong-password"

		// Action
		hash, err := auth.GeneratePasswordHash(password)

		// Assert
		assert.NotEqual(t, password, hash)
		assert.NoError(t, err)
	})
}

func TestCheckPasswordHash(t *testing.T) {
	t.Parallel()

	password := "my-secret-password"
	hash, err := auth.GeneratePasswordHash(password)
	assert.NoError(t, err)

	tests := []struct {
		about    string
		hash     string
		password string
		isValid  bool
	}{
		{
			about:    "when password is correct",
			hash:     hash,
			password: password,
			isValid:  true,
		},
		{
			about:    "when password is incorrect",
			hash:     hash,
			password: "another-password",
			isValid:  false,
		},
		{
			about:    "when hash is invalid",
			hash:     "",
			password: password,
			isValid:  false,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			ok := auth.CheckPasswordHash(tc.password, tc.hash)

			assert.Equal(t, tc.isValid, ok)
		})
	}
}
