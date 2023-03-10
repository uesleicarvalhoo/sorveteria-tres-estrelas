package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/payment"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/trace"
)

func Payments(router fiber.Router, service payment.UseCase) {
	router.Get("/", getPayments(service))
	router.Post("/", createPayment(service))
	router.Delete("/:id", deletePaymentByID(service))
	router.Post("/:id", updatePayment(service))
}

// @Summary      List payments
// @Description  get payments
// @Tags         Payment
// @Accept       json
// @Produce      json
// @Param        start_at    query   string  false  "name search by q"  Format(dateTime)
// @Param        end_at      query   string  false  "name search by q"  Format(dateTime)
// @Success      200         {object} []payment.Payment
// @Failure      500         {object} dto.MessageJSON "when an error occurs"
// @Router       /payments [get]
func getPayments(svc payment.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "get-payments")
		defer span.End()

		var query dto.GetPaymentByPeriodQuery

		if err := c.QueryParser(&query); err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusBadRequest).JSON(dto.MessageJSON{Message: err.Error()})
		}

		if query.StartAt.IsZero() || query.EndAt.IsZero() {
			payments, err := svc.GetAll(ctx)
			if err != nil {
				trace.AddSpanError(span, err)

				return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
			}

			return c.JSON(payments)
		}

		payments, err := svc.GetByPeriod(ctx, query.StartAt, query.EndAt)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.JSON(payments)
	}
}

// @Summary     Register a new payment
// @Description Create a new payment and return payment data
// @Tags        Payment
// @Accept      json
// @Produce     json
// @Param       payload body     dto.CreatePaymentPayload true "the payload data"
// @Success     201     {object} payment.Payment
// @Failure     422     {object} dto.MessageJSON "when payload is invalid"
// @Failure     500     {object} dto.MessageJSON "when an error occurs"
// @Router      /payments [post]
func createPayment(svc payment.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "register-payment")
		defer span.End()

		var payload payment.Payment

		if err := c.BodyParser(&payload); err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		p, err := svc.RegisterPayment(ctx, payload.Value, payload.Description)
		if err != nil {
			trace.AddSpanError(span, err)

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
func deletePaymentByID(svc payment.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "delete-payment")

		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		trace.AddSpanTags(span, map[string]string{"payment-id": id.String()})

		err = svc.DeletePayment(ctx, id)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.SendStatus(fiber.StatusAccepted)
	}
}

// @Summary     Update Payment by ID
// @Description Update payment data
// @Tags        Payment
// @Accept      json
// @Produce     json
// @Param       id path string true "the id of payment"
// @Success     200 {object} payment.Payment
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /payments/{id} [post]
func updatePayment(svc payment.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "update-payment")
		defer span.End()

		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		var payload dto.UpdatePaymentPayload
		if err := c.BodyParser(&payload); err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusBadRequest).JSON(dto.MessageJSON{Message: err.Error()})
		}

		trace.AddSpanTags(span, map[string]string{"payment-id": id.String()})

		p, err := svc.UpdatePayment(ctx, id, payload.Value, payload.Description)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(p)
	}
}
