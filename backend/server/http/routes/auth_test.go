//go:build unit || all

package routes_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/auth/mocks"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/server/http/routes"
)

func TestLogin(t *testing.T) {
	t.Parallel()

	t.Run("test success", func(t *testing.T) {
		t.Parallel()

		// Arrange
		payload := auth.LoginPayload{
			Email:    "user@email.com",
			Password: "secret123",
		}

		jwtToken := auth.JwtToken{
			GrantType: "beaerer",
			Token:     "my-access-token",
			ExpiresAt: time.Now().Unix(),
		}

		svc := mocks.NewUseCase(t)
		svc.On("Login", mock.Anything, payload).Return(jwtToken, nil).Once()

		app := fiber.New()
		routes.Auth(app, svc)

		reqBody, err := json.Marshal(payload)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
		req.Header.Add("Content-Type", "application/json")

		// Action
		res, err := app.Test(req)
		if assert.NoError(t, err) {
			defer res.Body.Close()
		}

		var body auth.JwtToken
		err = json.NewDecoder(res.Body).Decode(&body)
		assert.NoError(t, err)

		// Assert
		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.Equal(t, jwtToken.GrantType, body.GrantType)
		assert.Equal(t, jwtToken.Token, body.Token)
	})

	t.Run("test errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			about              string
			id                 string
			payload            auth.LoginPayload
			mockReturn         auth.JwtToken
			mockError          error
			expectedStatusCode int
			expectedBody       map[string]any
		}{
			{
				about:     "when service return an error",
				mockError: errors.New("service error"),
				payload: auth.LoginPayload{
					Email:    "user@email.com",
					Password: "secret123",
				},
				expectedStatusCode: http.StatusInternalServerError,
				expectedBody:       map[string]any{"message": "service error"},
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()

				// Arrange
				svc := mocks.NewUseCase(t)
				svc.On("Login", mock.Anything, tc.payload).
					Return(tc.mockReturn, tc.mockError).Once()

				app := fiber.New()
				routes.Auth(app, svc)

				reqBody, err := json.Marshal(tc.payload)
				assert.NoError(t, err)

				req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
				req.Header.Add("Content-Type", "application/json")

				// Action
				res, err := app.Test(req)
				if assert.NoError(t, err) {
					defer res.Body.Close()
				}

				var body map[string]any
				err = json.NewDecoder(res.Body).Decode(&body)
				assert.NoError(t, err)

				// Assert
				assert.Equal(t, tc.expectedStatusCode, res.StatusCode)
				assert.Equal(t, tc.expectedBody, body)
			})
		}
	})
}
