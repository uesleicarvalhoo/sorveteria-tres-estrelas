//go:build unit || all

package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/api/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/api/handler"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/sales/mocks"
)

func TestRegisterSale(t *testing.T) {
	t.Parallel()

	t.Run("test success", func(t *testing.T) {
		t.Parallel()

		payload := dto.RegisterSalePayload{
			Description: "sale description",
			PaymentType: entity.CashPayment,
			Items:       []entity.CartItem{{PopsicleID: uuid.New(), Amount: 30}},
		}

		createdSale := entity.Sale{
			ID:          uuid.New(),
			PaymentType: entity.CashPayment,
			Items:       []entity.SaleItem{{Name: "picole de coco", UnitPrice: 1.0, Amount: payload.Items[0].Amount}},
			Date:        time.Now(),
			Total:       float64(payload.Items[0].Amount),
			Description: payload.Description,
		}

		// Arrange
		svc := mocks.NewUseCase(t)
		svc.On("RegisterSale", mock.Anything,
			payload.Description, payload.PaymentType, entity.Cart{Items: payload.Items}).
			Return(createdSale, nil).Once()

		app := fiber.New()

		handler.MakeSalesRoutes(app, svc)

		resBody, err := json.Marshal(payload)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(resBody))
		req.Header.Add("Content-Type", "application/json")

		// Action
		res, err := app.Test(req, 30)
		assert.NoError(t, err)
		defer res.Body.Close()

		var body entity.Sale
		err = json.NewDecoder(res.Body).Decode(&body)
		assert.NoError(t, err)

		// Assert
		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, createdSale.ID, body.ID)
	})

	t.Run("test errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			about              string
			payload            dto.RegisterSalePayload
			mockReturn         entity.Sale
			mockError          error
			expectedStatusCode int
			expectedBody       map[string]any
		}{
			{
				about:              "when service return an error",
				payload:            dto.RegisterSalePayload{},
				mockReturn:         entity.Sale{},
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
				svc.On("RegisterSale", mock.Anything,
					tc.payload.Description, tc.payload.PaymentType, entity.Cart{Items: tc.payload.Items}).
					Return(tc.mockReturn, tc.mockError).Once()

				app := fiber.New()

				handler.MakeSalesRoutes(app, svc)

				resBody, err := json.Marshal(tc.payload)
				assert.NoError(t, err)

				req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(resBody))
				req.Header.Add("Content-Type", "application/json")

				// Action
				res, err := app.Test(req, 30)
				assert.NoError(t, err)
				defer res.Body.Close()

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