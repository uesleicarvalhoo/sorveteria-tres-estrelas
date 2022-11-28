package handler

import (
	"github.com/gofiber/fiber/v2"
)

func MakeHealthCheckRoutes(r fiber.Router) {
	r.Get("/health", healthCheck())
}

// @Summary     Health Check
// @Description Check app and dependencies status
// @Tags        Health Check
// @Produce     json
// @Success     200 {object} map[string]string
// @Router      /health [get]
func healthCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	}
}
