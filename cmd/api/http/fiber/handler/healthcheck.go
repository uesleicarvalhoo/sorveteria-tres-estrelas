package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/healthcheck"
)

func MakeHealthCheckRoutes(r fiber.Router, svc healthcheck.UseCase) {
	r.Get("/health", healthCheck(svc))
}

// @Summary		Health Cehck
// @Description	Check app and dependencies status
// @Tags		Health check
// @Produce		json
// @Success		200	{object} healthcheck.HealthStatus
// @Router		/health [get]
func healthCheck(svc healthcheck.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		s := svc.HealthCheck(c.Context())

		if s.Status != healthcheck.StatusUp {
			return c.Status(http.StatusInternalServerError).JSON(s)
		}

		return c.Status(http.StatusOK).JSON(s)
	}
}
