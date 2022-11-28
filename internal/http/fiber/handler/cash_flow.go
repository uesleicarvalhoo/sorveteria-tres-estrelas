package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cashflow"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/http/dto"
)

func MakeCashFlowHandler(r fiber.Router, svc cashflow.UseCase) {
	r.Get("/", cashflowIndex(svc))
}

// @Summary      Get cash flow
// @Description  get cash flow
// @Tags         Cashflow
// @Accept       json
// @Produce      json
// @Param        start_at    query   string  false  "name search by q"  Format(dateTime)
// @Param        end_at      query   string  false  "name search by q"  Format(dateTime)
// @Success      200         {object} cashflow.CashFlow
// @Failure      400         {object} dto.MessageJSON "when query is invalid"
// @Failure      500         {object} dto.MessageJSON "when an error occurs"
// @Router       /cashflow [get]
func cashflowIndex(svc cashflow.UseCase) fiber.Handler { //nolint:dupl
	return func(c *fiber.Ctx) error {
		var query dto.GetCashFlowByPeriodQuery
		if err := c.QueryParser(&query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.MessageJSON{Message: err.Error()})
		}

		if query.StartAt.IsZero() || query.EndAt.IsZero() {
			cf, err := svc.GetCashFlow(c.Context())
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
			}

			return c.JSON(cf)
		}

		cf, err := svc.GetCashFlowBetween(c.Context(), query.StartAt, query.EndAt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.JSON(cf)
	}
}
