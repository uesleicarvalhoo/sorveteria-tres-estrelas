package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/http/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payments"
)

func MakePaymentsRoutes(router fiber.Router, service payments.UseCase) {
	router.Get("/", getPayments(service))
	router.Post("/", createPayment(service))
	router.Delete("/:id", deletePaymentByID(service))
}

// @Summary      List payments
// @Description  get payments
// @Tags         Payments
// @Accept       json
// @Produce      json
// @Param        start_at    query   string  false  "name search by q"  Format(dateTime)
// @Param        end_at      query   string  false  "name search by q"  Format(dateTime)
// @Success      200         {object} []payments.Payment
// @Failure      500         {object} dto.MessageJSON "when an error occurs"
// @Router       /payments [get]
func getPayments(svc payments.UseCase) fiber.Handler { //nolint:dupl
	return func(c *fiber.Ctx) error {
		var query dto.GetPaymentByPeriodQuery

		if err := c.QueryParser(&query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.MessageJSON{Message: err.Error()})
		}

		if query.StartAt.IsZero() || query.EndAt.IsZero() {
			payments, err := svc.GetAll(c.Context())
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
			}

			return c.JSON(payments)
		}

		payments, err := svc.GetByPeriod(c.Context(), query.StartAt, query.EndAt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.JSON(payments)
	}
}

// @Summary     Register a new payment
// @Description Create a new payment and return payment data
// @Tags        Payments
// @Accept      json
// @Produce     json
// @Param       payload body     dto.CreatePaymentPayload true "the payload data"
// @Success     201     {object} payments.Payment
// @Failure     400     {object} dto.MessageJSON "when query is invalid"
// @Failure     500     {object} dto.MessageJSON "when an error occurs"
// @Router      /payments [post]
func createPayment(svc payments.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload payments.Payment
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		p, err := svc.RegisterPayment(c.Context(), payload.Value, payload.Description)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(p)
	}
}

// @Summary     Delete Payment by ID
// @Description Delete payment
// @Tags        Payment
// @Accept      json
// @Produce     json
// @Param       id path string true "the id of payment"
// @Success     202
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /payments/{id} [delete]
func deletePaymentByID(svc payments.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		err = svc.DeletePayment(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.SendStatus(fiber.StatusAccepted)
	}
}
