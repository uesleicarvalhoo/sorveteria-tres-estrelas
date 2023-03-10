package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/trace"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/user"
)

const UserIDHeader = "X-Kong-Jwt-Claim-Sub"

func User(r fiber.Router, userSvc user.UseCase) {
	r.Post("/", createUser(userSvc))
	r.Get("/", getMe(userSvc))
}

// @Summary     Create User
// @Description Create a new user and return user data
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       payload body dto.CreateUserPayload true "the user data"
// @Success     201 {object} user.User
// @Failure     422 {object} dto.MessageJSON "when payload is invalid"
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /users [post]
func createUser(svc user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "create-user")
		defer span.End()

		var payload dto.CreateUserPayload

		if err := c.BodyParser(&payload); err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		user, err := svc.Create(ctx, payload.Name, payload.Email, payload.Password)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(user)
	}
}

// @Summary     Get Me
// @Description Get current user data
// @Tags        User
// @Accept      json
// @Produce     json
// @Success     200 {object} user.User
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /user/me [get]
func getMe(svc user.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "get-me")
		defer span.End()

		sub := c.Get(UserIDHeader)

		userID, err := uuid.Parse(sub)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.MessageJSON{Message: "user not found"})
		}

		u, err := svc.Get(ctx, userID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.JSON(u)
	}
}
