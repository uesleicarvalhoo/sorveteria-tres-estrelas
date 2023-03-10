package middleware

import (
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
)

func Otel(serverName string) fiber.Handler {
	m := otelfiber.Middleware(serverName)

	return func(c *fiber.Ctx) error {
		if c.Path() == "/health" {
			return c.Next()
		}

		return m(c)
	}
}
