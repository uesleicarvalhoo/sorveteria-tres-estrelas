//go:build unit || all

package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/api/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/api/handler"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/popsicle/mocks"
)

func TestPopsicleGet(t *testing.T) {
	t.Parallel()

	t.Run("test success", func(t *testing.T) {
		t.Parallel()

		// Arrange
		storedPopsicle := entity.Popsicle{ID: uuid.New(), Flavor: "limão", Price: 1.25}

		svc := mocks.NewUseCase(t)
		svc.On("Get", mock.Anything, storedPopsicle.ID).Return(storedPopsicle, nil).Once()

		app := fiber.New()
		handler.MakePopsicleRoutes(app, svc)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", storedPopsicle.ID.String()), nil)

		// Action
		res, err := app.Test(req)
		assert.NoError(t, err)
		defer res.Body.Close()

		body := entity.Popsicle{}
		err = json.NewDecoder(res.Body).Decode(&body)
		assert.NoError(t, err)

		// Assert
		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.Equal(t, storedPopsicle, body)
	})

	t.Run("test errors", func(t *testing.T) {
		t.Parallel()

		storedPopsicle := entity.Popsicle{ID: uuid.New(), Flavor: "limão", Price: 1.25}

		tests := []struct {
			about              string
			id                 string
			mockReturn         entity.Popsicle
			mockError          error
			expectedStatusCode int
			expectedBody       map[string]any
		}{
			{
				about:              "when service return an error",
				id:                 storedPopsicle.ID.String(),
				mockError:          errors.New("service error"),
				expectedStatusCode: http.StatusInternalServerError,
				expectedBody:       map[string]any{"message": "service error"},
			},
			{
				about:              "when id is invalid",
				id:                 "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				expectedStatusCode: http.StatusUnprocessableEntity,
				expectedBody:       map[string]any{"message": "invalid UUID format"},
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()

				// Arrange
				svc := mocks.NewUseCase(t)
				svc.On("Get", mock.Anything, storedPopsicle.ID).Return(tc.mockReturn, tc.mockError).Maybe()

				app := fiber.New()
				handler.MakePopsicleRoutes(app, svc)

				req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", tc.id), nil)

				// Action
				res, err := app.Test(req)
				assert.NoError(t, err)
				defer res.Body.Close()

				body := map[string]any{}
				err = json.NewDecoder(res.Body).Decode(&body)
				assert.NoError(t, err)

				// Assert
				assert.Equal(t, res.StatusCode, tc.expectedStatusCode)
				assert.Equal(t, tc.expectedBody, body)
			})
		}
	})
}

func TestPopsicleIndex(t *testing.T) {
	t.Parallel()
	id := uuid.MustParse("d64edd2f-013a-43b5-8c68-9675bbf0e840")

	t.Run("test success", func(t *testing.T) {
		t.Parallel()

		// Arrange
		storedPopsicles := []entity.Popsicle{
			{ID: id, Flavor: "limão", Price: 1.25},
		}

		svc := mocks.NewUseCase(t)
		svc.On("Index", mock.Anything).Return(storedPopsicles, nil).Once()

		app := fiber.New()
		handler.MakePopsicleRoutes(app, svc)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		// Action
		res, err := app.Test(req)
		defer res.Body.Close()

		assert.NoError(t, err)

		var body []entity.Popsicle
		json.NewDecoder(res.Body).Decode(&body)
		assert.NoError(t, err)

		// Assert
		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.Equal(t, storedPopsicles, body)
	})

	t.Run("test errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			about              string
			mockReturn         []entity.Popsicle
			mockError          error
			expectedStatusCode int
			expectedBody       []byte
		}{
			{
				about:              "when service return a empty list",
				mockReturn:         []entity.Popsicle{},
				expectedStatusCode: http.StatusOK,
				expectedBody:       []byte("[]"),
			},
			{
				about:              "when service return an error",
				mockError:          errors.New("service error"),
				expectedStatusCode: http.StatusInternalServerError,
				expectedBody:       []byte("{\"message\":\"service error\"}"),
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()

				// Arrange
				svc := mocks.NewUseCase(t)
				svc.On("Index", mock.Anything).Return(tc.mockReturn, tc.mockError).Once()

				app := fiber.New()
				handler.MakePopsicleRoutes(app, svc)

				req := httptest.NewRequest(http.MethodGet, "/", nil)

				// Action
				res, err := app.Test(req)
				defer res.Body.Close()

				assert.NoError(t, err)

				body, err := io.ReadAll(res.Body)
				assert.NoError(t, err)

				// Assert
				assert.Equal(t, res.StatusCode, tc.expectedStatusCode)
				assert.Equal(t, tc.expectedBody, body)
			})
		}
	})
}

func TestPopsicleStore(t *testing.T) {
	t.Run("test success", func(t *testing.T) {
		t.Parallel()

		// Arrange
		payload := dto.CreatePopsiclePayload{Flavor: "amendoin", Price: 1.23}
		createdPopsicle := entity.Popsicle{ID: uuid.New(), Flavor: payload.Flavor, Price: payload.Price}

		svc := mocks.NewUseCase(t)
		svc.On("Store", mock.Anything, payload.Flavor, payload.Price).Return(createdPopsicle, nil).Once()

		app := fiber.New()

		handler.MakePopsicleRoutes(app, svc)

		reqBody, err := json.Marshal(payload)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		req.Header.Add("Content-Type", "application/json")

		// Action
		res, err := app.Test(req, 30)
		assert.NoError(t, err)
		defer res.Body.Close()

		var body entity.Popsicle
		err = json.NewDecoder(res.Body).Decode(&body)
		assert.NoError(t, err)

		// Assert
		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, createdPopsicle, body)
	})

	t.Run("test errors", func(t *testing.T) {
		t.Parallel()

		createdPopsicle := entity.Popsicle{
			ID:     uuid.New(),
			Flavor: "amendoin",
			Price:  1.23,
		}

		tests := []struct {
			about              string
			payload            dto.CreatePopsiclePayload
			mockReturn         entity.Popsicle
			mockError          error
			expectedStatusCode int
			expectedBody       map[string]any
		}{
			{
				about:              "when service return an error",
				payload:            dto.CreatePopsiclePayload{Flavor: createdPopsicle.Flavor, Price: createdPopsicle.Price},
				mockReturn:         entity.Popsicle{},
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
				svc.On("Store", mock.Anything, tc.payload.Flavor, tc.payload.Price).Return(tc.mockReturn, tc.mockError).Once()

				app := fiber.New()

				handler.MakePopsicleRoutes(app, svc)

				payload, err := json.Marshal(tc.payload)
				assert.NoError(t, err)

				req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
				req.Header.Add("Content-Type", "application/json")

				// Action
				res, err := app.Test(req, 30)
				assert.NoError(t, err)
				defer res.Body.Close()

				body := map[string]any{}
				err = json.NewDecoder(res.Body).Decode(&body)
				assert.NoError(t, err)

				// Assert
				assert.Equal(t, tc.expectedStatusCode, res.StatusCode)
				assert.Equal(t, tc.expectedBody, body)
			})
		}
	})
}
