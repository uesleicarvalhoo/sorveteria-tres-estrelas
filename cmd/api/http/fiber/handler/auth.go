package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/trace"
)

func MakeAuhtRoutes(r fiber.Router, authSvc auth.UseCase) {
	r.Post("/login", login(authSvc))
	r.Post("refresh-token", refreshToken(authSvc))
}

// @Summary     Login
// @Description Make login and get access-token
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       payload      body dto.LoginPayload true "user info"
// @Success     200 {object} auth.JwtToken
// @Failure     401 {object} dto.MessageJSON "when email or password is invalid"
// @Failure     422 {object} dto.MessageJSON "when payload is invalid"
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /auth/login [post]
func login(svc auth.UseCase) fiber.Handler { //nolint:dupl
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "login")
		defer span.End()

		var payload dto.LoginPayload

		if err := c.BodyParser(&payload); err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		token, err := svc.Login(ctx, payload.Email, payload.Password)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(token)
	}
}

// @Summary Refresh access-token
// @Description Get a new access-token, this action will be expire the last one
// @Tags    Auth
// @Accept  json
// @Produce json
// @Param   payload	body dto.RefreshTokenPayload true "the refresh token"
// @Success 200 {object} auth.JwtToken
// @Failure 401 {object} dto.MessageJSON "when token is invalid"
// @Failure 500 {object} dto.MessageJSON "when an error occurs"
// @Router  /auth/refresh-token [post]
func refreshToken(svc auth.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "refresh-token")
		defer span.End()

		var payload dto.RefreshTokenPayload

		if err := c.BodyParser(&payload); err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		token, err := svc.RefreshToken(ctx, payload.RefreshToken)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(token)
	}
}
