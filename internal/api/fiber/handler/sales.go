package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
)

func MakeSalesRoutes(r fiber.Router, svc sales.UseCase, balanceSvc balances.UseCase) {
	r.Get("/", salesIndex(svc))
	r.Post("/", registerSale(svc, balanceSvc))
}

// @Summary      List sales
// @Description  get sales
// @Tags         Sales
// @Accept       json
// @Produce      json
// @Param        startAt    query   string  false  "name search by q"  Format(dateTime)
// @Param        endAt      query   string  false  "name search by q"  Format(dateTime)
// @Success		200	{object} []sales.Sale
// @Failure		422	{object} dto.MessageJSON "when start or end param is invalid"
// @Failure		500	{object} dto.MessageJSON "when an error occurs"
// @Router		/sales [get]
func salesIndex(svc sales.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload dto.GetSalesByPeriodQuery
		if err := c.QueryParser(&payload); err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		if payload.EndAt.IsZero() || payload.StartAt.IsZero() {
			sales, err := svc.GetAll(c.Context())
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
			}

			return c.JSON(sales)
		}

		sales, err := svc.GetByPeriod(c.Context(), payload.StartAt, payload.EndAt)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.JSON(sales)
	}
}

// @Summary		Register a new sale
// @Description	Register a sale and return sale data
// @Tags		Sales
// @Accept		json
// @Produce		json
// @Param		payload		body	dto.RegisterSalePayload true "the payload data"
// @Success		200	{object} []sales.Sale
// @Failure		422	{object} dto.MessageJSON "when payload is invalid"
// @Failure		500	{object} dto.MessageJSON "when an error occurs"
// @Router		/sales [post]
func registerSale(svc sales.UseCase, balanceSvc balances.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload dto.RegisterSalePayload

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		sale, err := svc.RegisterSale(
			c.Context(), payload.Description, payload.PaymentType, sales.Cart{Items: payload.Items})
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		if _, err := balanceSvc.RegisterFromSale(c.Context(), sale); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.Status(http.StatusCreated).JSON(sale)
	}
}
