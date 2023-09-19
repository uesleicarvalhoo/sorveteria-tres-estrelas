package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/trace"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/transaction"
)

func Transacation(router fiber.Router, service transaction.UseCase) {
	router.Get("/", getTransactions(service))
	router.Post("/", createTransaction(service))
	router.Delete("/:id", deleteTransactionByID(service))
}

// @Summary      List transactions
// @Description  get transactions
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Param        start_at    query   string  false  "name search by q"  Format(dateTime)
// @Param        end_at      query   string  false  "name search by q"  Format(dateTime)
// @Success      200         {object} []transaction.Transaction
// @Failure      500         {object} dto.MessageJSON "when an error occurs"
// @Router       /transactions [get]
func getTransactions(svc transaction.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "get-transactions")
		defer span.End()

		var query dto.GetTransactionByPeriodQuery

		if err := c.QueryParser(&query); err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusBadRequest).JSON(dto.MessageJSON{Message: err.Error()})
		}

		if query.StartAt.IsZero() || query.EndAt.IsZero() {
			transactions, err := svc.GetTransactions(ctx)
			if err != nil {
				trace.AddSpanError(span, err)

				return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
			}

			return c.JSON(transactions)
		}

		transactions, err := svc.GetByPeriod(ctx, query.StartAt, query.EndAt)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.JSON(transactions)
	}
}

// @Summary     Register a new transaction
// @Description Create a new transaction and return transaction data
// @Tags        Transaction
// @Accept      json
// @Produce     json
// @Param       payload body     dto.CreateTransactionPayload true "the payload data"
// @Success     201     {object} transaction.Transaction
// @Failure     422     {object} dto.MessageJSON "when payload is invalid"
// @Failure     500     {object} dto.MessageJSON "when an error occurs"
// @Router      /transactions [post]
func createTransaction(svc transaction.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "register-payment")
		defer span.End()

		var payload transaction.Transaction

		if err := c.BodyParser(&payload); err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		p, err := svc.RegisterTransaction(ctx, payload.Value, payload.Type, payload.Description)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(p)
	}
}

// @Summary     Delete Transaction by ID
// @Description Delete payment
// @Tags        Transaction
// @Accept      json
// @Produce     json
// @Param       id path string true "the id of payment"
// @Success     202
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /transactions/{id} [delete]
func deleteTransactionByID(svc transaction.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "delete-payment")
		defer span.End()

		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		trace.AddSpanTags(span, map[string]string{"payment-id": id.String()})

		err = svc.DeleteTransaction(ctx, id)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.SendStatus(fiber.StatusAccepted)
	}
}
