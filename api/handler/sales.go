package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/api/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/sales"
)

func MakeSalesRoutes(r fiber.Router, svc sales.UseCase) {
	r.Get("/", salesIndex(svc))
	r.Post("/", registerSale(svc))
}

// @Summary		List Sales
// @Description	Get all sales
// @Tags		Sales
// @Produce		json
// @Param		start	query	time.Time	false	"the start date of period to search sales"
// @Param		end		query	time.Time	false	"the end date of period to search sales"
// @Success		200	{object} []entity.Sale
// @Failure		422	{object} dto.MessageJSON "when start or end param is invalid"
// @Failure		500	{object} dto.MessageJSON "when an error occurs"
// @Router		/sales [get].
func salesIndex(svc sales.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload dto.GetSalesByPeriodQuery
		if err := c.QueryParser(&payload); err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
		}

		if payload.EndAt.IsZero() || payload.StartAt.IsZero() {
			sales, err := svc.GetAll(c.Context())
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
			}

			return c.JSON(sales)
		}

		sales, err := svc.GetByPeriod(c.Context(), payload.StartAt, payload.EndAt)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
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
// @Success		200	{object} []entity.Sale
// @Failure		422	{object} dto.MessageJSON "when payload is invalid"
// @Failure		500	{object} dto.MessageJSON "when an error occurs"
// @Router		/sales [post].
func registerSale(svc sales.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload dto.RegisterSalePayload

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
		}

		sale, err := svc.RegisterSale(
			c.Context(), payload.Description, payload.PaymentType, entity.Cart{Items: payload.Items})
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}

		return c.Status(http.StatusCreated).JSON(sale)
	}
}
