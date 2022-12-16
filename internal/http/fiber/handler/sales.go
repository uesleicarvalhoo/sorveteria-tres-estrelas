package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/http/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
)

func MakeSalesRoutes(r fiber.Router, svc sales.UseCase) {
	r.Get("/", salesIndex(svc))
	r.Post("/", registerSale(svc))
	r.Delete("/:id", deleteSaleByID(svc))
}

// @Summary      List sales
// @Description  get sales
// @Tags         Sale
// @Accept       json
// @Produce      json
// @Param        start_at    query   string  false  "name search by q"  Format(dateTime)
// @Param        end_at      query   string  false  "name search by q"  Format(dateTime)
// @Success      200 {object} []sales.Sale
// @Failure      400 {object} dto.MessageJSON "when start or end param is invalid"
// @Failure      500 {object} dto.MessageJSON "when an error occurs"
// @Router       /sales [get]
func salesIndex(svc sales.UseCase) fiber.Handler { //nolint:dupl
	return func(c *fiber.Ctx) error {
		var payload dto.GetSalesByPeriodQuery
		if err := c.QueryParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.MessageJSON{Message: err.Error()})
		}

		if payload.EndAt.IsZero() || payload.StartAt.IsZero() {
			sales, err := svc.GetAll(c.Context())
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
			}

			return c.JSON(sales)
		}

		sales, err := svc.GetByPeriod(c.Context(), payload.StartAt, payload.EndAt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.JSON(sales)
	}
}

// @Summary     Register a new sale
// @Description Register a sale and return sale data
// @Tags        Sale
// @Accept      json
// @Produce     json
// @Param       payload body     dto.RegisterSalePayload true "the payload data"
// @Success     200     {object} []sales.Sale
// @Failure     422     {object} dto.MessageJSON "when payload is invalid"
// @Failure     500     {object} dto.MessageJSON "when an error occurs"
// @Router      /sales [post]
func registerSale(svc sales.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload dto.RegisterSalePayload

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		sale, err := svc.RegisterSale(
			c.Context(), payload.Description, payload.PaymentType, sales.Cart{Items: payload.Items})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(sale)
	}
}

// @Summary     Delete Sale by ID
// @Description Delete sale
// @Tags        Sale
// @Accept      json
// @Produce     json
// @Param       id path string true "the id of sale"
// @Success     202
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /sales/{id} [delete]
func deleteSaleByID(svc sales.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		err = svc.DeleteByID(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.SendStatus(fiber.StatusAccepted)
	}
}
