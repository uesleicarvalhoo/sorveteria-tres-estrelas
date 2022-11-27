package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/http/dto"
)

func MakeBalanceRouter(r fiber.Router, svc balances.UseCase) {
	r.Get("/", balancesIndex(svc))
	r.Post("/", createBalance(svc))
}

// @Summary      List balances
// @Description  get balances
// @Tags         Balances
// @Accept       json
// @Produce      json
// @Param        start_at    query   string  false  "name search by q"  Format(dateTime)
// @Param        end_at      query   string  false  "name search by q"  Format(dateTime)
// @Success      200         {object} []balances.CashFlow
// @Failure      500         {object} dto.MessageJSON "when an error occurs"
// @Router       /balances [get]
func balancesIndex(svc balances.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query dto.GetCashFlowByPeriodQuery
		if err := c.QueryParser(&query); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		if !query.StartAt.IsZero() && !query.EndAt.IsZero() {
			balances, err := svc.GetCashFlowBetween(c.Context(), query.StartAt, query.EndAt)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
			}

			return c.JSON(balances)
		}

		balances, err := svc.GetCashFlow(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.JSON(balances)
	}
}

// @Summary		Register a new balance
// @Description	Register a balance and return balance data
// @Tags			Balances
// @Accept		json
// @Produce		json
// @Param		payload		body	dto.RegisterBalancePayload true "the payload data"
// @Success		201	{object} balances.Balance
// @Failure		422	{object} dto.MessageJSON "when payload is invalid"
// @Failure		500	{object} dto.MessageJSON "when an error occurs"
// @Router		/balances [post]
func createBalance(svc balances.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload dto.RegisterBalancePayload

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{
				Message: err.Error(),
			})
		}

		balance, err := svc.RegisterOperation(c.Context(), payload.Value, payload.Description, payload.Operation)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(balance)
	}
}
