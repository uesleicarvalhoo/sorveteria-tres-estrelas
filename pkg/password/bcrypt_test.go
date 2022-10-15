//go:build unit || all

package password_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/pkg/password"
)

func TestGenerateHash(t *testing.T) {
	t.Parallel()

	t.Run("invalid password", func(t *testing.T) {
		t.Parallel()

		// Arrange
		sut := password.NewBCrypt()
		passwd := ""

		// Action
		hash, err := sut.GenerateHash(passwd)

		// Assert
		assert.Equal(t, "", hash)
		assert.EqualError(t, err, password.ErrTooShortPassword.Error())
	})

	t.Run("valid password", func(t *testing.T) {
		t.Parallel()

		// Arrange
		sut := password.NewBCrypt()

		passwd := "strong-password"

		// Action
		hash, err := sut.GenerateHash(passwd)

		// Assert
		assert.NotEqual(t, passwd, hash)
		assert.NoError(t, err)
	})
}

func TestCheckHash(t *testing.T) {
	t.Parallel()

	passwd := "my-secret-password"

	sut := password.NewBCrypt()

	hash, err := sut.GenerateHash(passwd)
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
			password: passwd,
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
			password: passwd,
			isValid:  false,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			ok := sut.CheckHash(tc.password, tc.hash)

			assert.Equal(t, tc.isValid, ok)
		})
	}
}
