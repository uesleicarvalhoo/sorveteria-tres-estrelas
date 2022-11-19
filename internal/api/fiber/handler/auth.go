package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api/dto"
)

func MakeAuhtRoutes(r fiber.Router, authSvc auth.UseCase) {
	r.Post("/login", login(authSvc))
	r.Post("refresh-token", refreshToken(authSvc))
}

// @Summary		Login
// @Description	Make login and get access-token
// @Tags		Auth
// @Accept      json
// @Produce     json
// @Param		payload		body	dto.LoginPayload	true	"user info"
// @Success		200	{object} []auth.JwtToken
// @Failure		401	{object} dto.MessageJSON "when email or password is invalid"
// @Failure		500	{object} dto.MessageJSON "when an error occurs"
// @Router		/auth/login [post]
func login(svc auth.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload dto.LoginPayload

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
		}

		token, err := svc.Login(c.Context(), payload.Email, payload.Password)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}

		return c.Status(http.StatusCreated).JSON(token)
	}
}

// @Summary		Refresh access-token
// @Description	Get a new access-token, this action will be expire the last one
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		payload	body	dto.RefreshTokenPayload	true	"the refresh token"
// @Success		200	{object} []auth.JwtToken
// @Failure		401	{object} dto.MessageJSON "when token is invalid"
// @Failure		500	{object} dto.MessageJSON "when an error occurs"
// @Router		/auth/refresh-token [post].
func refreshToken(svc auth.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload dto.RefreshTokenPayload

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
		}

		token, err := svc.RefreshToken(c.Context(), payload.RefreshToken)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}

		return c.Status(http.StatusCreated).JSON(token)
	}
}
