//go:build unit || all

package routes_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func mockAuthMiddleware(userID uuid.UUID) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("userID", userID)

		return c.Next()
	}
}
