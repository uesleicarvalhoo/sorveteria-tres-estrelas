package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/trace"
)

func Auth(r fiber.Router, authSvc auth.UseCase) {
	r.Post("/login", login(authSvc))
	r.Post("/refresh-token", refreshToken(authSvc))
	r.Get("/me", getMe(authSvc))
}

// @Summary     Login
// @Description Make login and get access-token
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       payload      body auth.LoginPayload true "user info"
// @Success     200 {object} auth.JwtToken
// @Failure     401 {object} dto.MessageJSON "when email or password is invalid"
// @Failure     422 {object} dto.MessageJSON "when payload is invalid"
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /auth/login [post]
func login(svc auth.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "login")
		defer span.End()

		var payload auth.LoginPayload

		if err := c.BodyParser(&payload); err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		token, err := svc.Login(ctx, payload)
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
// @Param   payload	body auth.RefreshTokenPayload true "the refresh token"
// @Success 200 {object} auth.JwtToken
// @Failure 401 {object} dto.MessageJSON "when token is invalid"
// @Failure 500 {object} dto.MessageJSON "when an error occurs"
// @Router  /auth/refresh-token [post]
func refreshToken(svc auth.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "refresh-token")
		defer span.End()

		var payload auth.RefreshTokenPayload

		if err := c.BodyParser(&payload); err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		token, err := svc.RefreshToken(ctx, payload)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(token)
	}
}

// @Summary     Get Me
// @Description Get current user data
// @Tags        User
// @Accept      json
// @Produce     json
// @Success     200 {object} user.User
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /auth/me [get]
func getMe(svc auth.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "get-me")
		defer span.End()

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.MessageJSON{Message: "missing 'Authorization' header"})
		}

		token := authHeader[len("Bearer "):]

		u, err := svc.Authorize(ctx, token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.JSON(u)
	}
}
