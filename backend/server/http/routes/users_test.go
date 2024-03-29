//go:build unit || all

package routes_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/server/http/routes"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/user"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/user/mocks"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	t.Run("should create user", func(t *testing.T) {
		t.Parallel()

		// Arrange
		payload := dto.CreateUserPayload{
			Name:     "User Lastname",
			Email:    "user@email.com",
			Password: "secret123",
		}

		storedUser, err := user.NewUser(payload.Name, payload.Email, payload.Password)
		assert.NoError(t, err)

		svc := mocks.NewUseCase(t)
		svc.On("Create", mock.Anything, payload.Name, payload.Email, payload.Password).
			Return(storedUser, nil).Once()

		app := fiber.New()
		routes.User(app, svc)

		reqBody, err := json.Marshal(payload)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		req.Header.Add("Content-Type", "application/json")

		// Action
		res, err := app.Test(req, 30)
		if assert.NoError(t, err) {
			defer res.Body.Close()
		}

		var body user.User
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
			mockReturn         user.User
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
				svc := mocks.NewUseCase(t)
				svc.On("Create", mock.Anything, tc.payload.Name, tc.payload.Email, tc.payload.Password).
					Return(tc.mockReturn, tc.mockError).Once()

				app := fiber.New()
				routes.User(app, svc)

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

func TestGetMe(t *testing.T) {
	t.Parallel()

	existingUser, err := user.NewUser("username", "user@email.com.br", "123456")
	assert.NoError(t, err)

	tests := []struct {
		about              string
		headers            map[string]string
		userID             uuid.UUID
		mockReturn         user.User
		mockError          error
		expectedStatusCode int
		expectedBody       map[string]any
	}{
		{
			about:              "when request has a X-User-ID header with an invalid UUID",
			headers:            map[string]string{"X-User-ID": "0"},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       map[string]any{"message": "invalid user"},
		},
		{
			about:              "when request has a valid X-User-ID header",
			headers:            map[string]string{"X-User-ID": existingUser.ID.String()},
			userID:             existingUser.ID,
			mockReturn:         existingUser,
			expectedStatusCode: http.StatusOK,
			expectedBody: map[string]any{
				"id":    existingUser.ID.String(),
				"name":  existingUser.Name,
				"email": existingUser.Email,
			},
		},
		{
			about:              "when service return an error",
			headers:            map[string]string{"X-User-ID": existingUser.ID.String()},
			userID:             existingUser.ID,
			mockReturn:         user.User{},
			mockError:          errors.New("service error"),
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
			svc.On("Get", mock.Anything, tc.userID).Return(tc.mockReturn, tc.mockError).Maybe()

			app := fiber.New()
			routes.User(app, svc)

			req := httptest.NewRequest(http.MethodGet, "/me", nil)
			for k, v := range tc.headers {
				req.Header.Add(k, v)
			}

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
}
