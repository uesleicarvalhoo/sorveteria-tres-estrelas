//go:build unit || all

package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api/fiber/handler"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users/mocks"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	t.Run("should create user", func(t *testing.T) {
		t.Parallel()

		// Arrange
		currentUser, err := users.NewUser(
			"current user", "current.user@email.com", "passwd")
		assert.NoError(t, err)

		payload := dto.CreateUserPayload{
			Name:     "User Lastname",
			Email:    "user@email.com",
			Password: "secret123",
		}

		storedUser, err := users.NewUser(payload.Name, payload.Email, payload.Password)
		assert.NoError(t, err)

		svc := mocks.NewUseCase(t)
		svc.On("Get", mock.Anything, currentUser.ID).Return(currentUser, nil)
		svc.On("Create", mock.Anything, payload.Name, payload.Email, payload.Password).
			Return(storedUser, nil).Once()

		app := fiber.New()
		app.Use(mockAuthMiddleware(currentUser.ID))

		handler.MakeUserRoutes(app, svc)

		reqBody, err := json.Marshal(payload)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		req.Header.Add("Content-Type", "application/json")

		// Action
		res, err := app.Test(req, 30)
		if assert.NoError(t, err) {
			defer res.Body.Close()
		}

		var body users.User
		err = json.NewDecoder(res.Body).Decode(&body)
		assert.NoError(t, err)

		// Assert
		assert.Equal(t, res.StatusCode, http.StatusCreated)
		assert.Equal(t, storedUser.ID, body.ID)
		assert.Equal(t, storedUser.Email, body.Email)
		assert.Equal(t, storedUser.Name, body.Name)
		assert.Equal(t, "", body.PasswordHash)
	})

	t.Run("test errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			about              string
			id                 string
			payload            dto.CreateUserPayload
			mockReturn         users.User
			mockError          error
			expectedStatusCode int
			expectedBody       map[string]any
		}{
			{
				about:     "when service return an error",
				mockError: errors.New("service error"),
				payload: dto.CreateUserPayload{
					Name:     "User Lastname",
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
				currentUser, err := users.NewUser(
					"current user", "current.user@email.com", "passwd")
				assert.NoError(t, err)

				svc := mocks.NewUseCase(t)
				svc.On("Create", mock.Anything, tc.payload.Name, tc.payload.Email, tc.payload.Password).
					Return(tc.mockReturn, tc.mockError).Once()

				app := fiber.New()
				app.Use(mockAuthMiddleware(currentUser.ID))
				handler.MakeUserRoutes(app, svc)

				reqBody, err := json.Marshal(tc.payload)
				assert.NoError(t, err)

				req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
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
