package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/api/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
)

func NewAuth(authSvc auth.UseCase, domain string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(dto.MessageJSON{Message: "authorization not found"})
		}

		token := authHeader[len("Bearer")+1:]

		perm := "read"
		if c.Method() != "GET" {
			perm = "write"
		}

		userID, err := authSvc.Authorize(c.Context(), token, domain, perm)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(dto.MessageJSON{Message: err.Error()})
		}

		c.Set("x-user-id", userID.String())
		c.Locals("userID", userID)

		return c.Next()
	}
}
