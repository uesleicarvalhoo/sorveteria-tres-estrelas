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
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cmd/api/http/fiber/handler"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/product"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/product/mocks"
)

func TestProductsGet(t *testing.T) {
	t.Parallel()

	t.Run("test success", func(t *testing.T) {
		t.Parallel()

		// Arrange
		storedProduct := product.Product{
			ID:            uuid.New(),
			Name:          "picole de limão",
			PriceVarejo:   1.25,
			PriceAtacado:  1,
			AtacadoAmount: 10,
		}

		svc := mocks.NewUseCase(t)
		svc.On("Get", mock.Anything, storedProduct.ID).Return(storedProduct, nil).Once()

		app := fiber.New()
		handler.MakeProductsRoutes(app, svc)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", storedProduct.ID.String()), nil)

		// Action
		res, err := app.Test(req)
		if assert.NoError(t, err) {
			defer res.Body.Close()
		}

		body := product.Product{}
		err = json.NewDecoder(res.Body).Decode(&body)
		assert.NoError(t, err)

		// Assert
		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.Equal(t, storedProduct, body)
	})

	t.Run("test errors", func(t *testing.T) {
		t.Parallel()

		storedProducts := product.Product{
			ID:            uuid.New(),
			Name:          "picole de limão",
			PriceVarejo:   1.25,
			PriceAtacado:  1,
			AtacadoAmount: 10,
		}

		tests := []struct {
			about              string
			id                 string
			mockReturn         product.Product
			mockError          error
			expectedStatusCode int
			expectedBody       map[string]any
		}{
			{
				about:              "when service return an error",
				id:                 storedProducts.ID.String(),
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
				svc.On("Get", mock.Anything, storedProducts.ID).Return(tc.mockReturn, tc.mockError).Maybe()

				app := fiber.New()
				handler.MakeProductsRoutes(app, svc)

				req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", tc.id), nil)

				// Action
				res, err := app.Test(req)
				if assert.NoError(t, err) {
					defer res.Body.Close()
				}

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

func TestProductsIndex(t *testing.T) {
	t.Parallel()
	id := uuid.MustParse("d64edd2f-013a-43b5-8c68-9675bbf0e840")

	t.Run("test success", func(t *testing.T) {
		t.Parallel()

		// Arrange
		storedProducts := []product.Product{
			{
				ID:            id,
				Name:          "picole de limão",
				PriceVarejo:   1.25,
				PriceAtacado:  1,
				AtacadoAmount: 10,
			},
		}

		svc := mocks.NewUseCase(t)
		svc.On("Index", mock.Anything).Return(storedProducts, nil).Once()

		app := fiber.New()
		handler.MakeProductsRoutes(app, svc)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		// Action
		res, err := app.Test(req)
		if assert.NoError(t, err) {
			defer res.Body.Close()
		}

		var body []product.Product
		json.NewDecoder(res.Body).Decode(&body)
		assert.NoError(t, err)

		// Assert
		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.Equal(t, storedProducts, body)
	})

	t.Run("test errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			about              string
			mockReturn         []product.Product
			mockError          error
			expectedStatusCode int
			expectedBody       []byte
		}{
			{
				about:              "when service return a empty list",
				mockReturn:         []product.Product{},
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
				handler.MakeProductsRoutes(app, svc)

				req := httptest.NewRequest(http.MethodGet, "/", nil)

				// Action
				res, err := app.Test(req)
				if assert.NoError(t, err) {
					defer res.Body.Close()
				}

				body, err := io.ReadAll(res.Body)
				assert.NoError(t, err)

				// Assert
				assert.Equal(t, res.StatusCode, tc.expectedStatusCode)
				assert.Equal(t, tc.expectedBody, body)
			})
		}
	})
}

func TestProductsStore(t *testing.T) {
	t.Run("test success", func(t *testing.T) {
		t.Parallel()

		// Arrange
		payload := dto.CreateProductPayload{
			Name: "picole de amendoin", PriceVarejo: 1.23, PriceAtacado: 1, AtacadoAmount: 10,
		}

		createdProduct := product.Product{
			ID:            uuid.New(),
			Name:          payload.Name,
			PriceVarejo:   payload.PriceVarejo,
			PriceAtacado:  payload.PriceAtacado,
			AtacadoAmount: int(payload.AtacadoAmount),
		}

		svc := mocks.NewUseCase(t)
		svc.On("Store", mock.Anything,
			payload.Name, payload.PriceVarejo, payload.PriceAtacado, payload.AtacadoAmount).
			Return(createdProduct, nil).Once()

		app := fiber.New()

		handler.MakeProductsRoutes(app, svc)

		reqBody, err := json.Marshal(payload)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		req.Header.Add("Content-Type", "application/json")

		// Action
		res, err := app.Test(req, 30)
		if assert.NoError(t, err) {
			defer res.Body.Close()
		}

		var body product.Product
		err = json.NewDecoder(res.Body).Decode(&body)
		assert.NoError(t, err)

		// Assert
		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, createdProduct, body)
	})

	t.Run("test errors", func(t *testing.T) {
		t.Parallel()

		createdProduct := product.Product{
			ID:            uuid.New(),
			Name:          "picole de limão",
			PriceVarejo:   1.25,
			PriceAtacado:  1,
			AtacadoAmount: 10,
		}

		tests := []struct {
			about              string
			payload            dto.CreateProductPayload
			mockReturn         product.Product
			mockError          error
			expectedStatusCode int
			expectedBody       map[string]any
		}{
			{
				about: "when service return an error",
				payload: dto.CreateProductPayload{
					Name:          createdProduct.Name,
					PriceVarejo:   createdProduct.PriceVarejo,
					PriceAtacado:  createdProduct.PriceAtacado,
					AtacadoAmount: createdProduct.AtacadoAmount,
				},
				mockReturn:         product.Product{},
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
				svc.On("Store", mock.Anything,
					tc.payload.Name, tc.payload.PriceVarejo, tc.payload.PriceAtacado, tc.payload.AtacadoAmount).
					Return(tc.mockReturn, tc.mockError).Once()

				app := fiber.New()

				handler.MakeProductsRoutes(app, svc)

				payload, err := json.Marshal(tc.payload)
				assert.NoError(t, err)

				req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
				req.Header.Add("Content-Type", "application/json")

				// Action
				res, err := app.Test(req, 30)
				if assert.NoError(t, err) {
					defer res.Body.Close()
				}

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
