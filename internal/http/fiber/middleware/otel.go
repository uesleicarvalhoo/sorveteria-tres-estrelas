package middleware

import (
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
)

func NewOtel(serverName string) fiber.Handler {
	return otelfiber.Middleware(serverName)
}
