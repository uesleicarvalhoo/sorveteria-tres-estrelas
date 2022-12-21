package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/http/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
)

func MakeUserRoutes(r fiber.Router, userSvc users.UseCase) {
	r.Get("/me", getMe(userSvc))
	r.Post("/", createUser(userSvc))
}

// @Summary     Get Me
// @Description Get current user data
// @Tags        User
// @Accept      json
// @Produce     json
// @Success     200 {object} users.User
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /users/me [get]
func getMe(svc users.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u, _ := c.Locals("user").(*users.User)
		if u == nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: "user not found"})
		}

		return c.JSON(u)
	}
}

// @Summary     Create User
// @Description Create a new user and return user data
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       payload body dto.CreateUserPayload true "the user data"
// @Success     201 {object} users.User
// @Failure     422 {object} dto.MessageJSON "when payload is invalid"
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /users [post]
func createUser(svc users.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload dto.CreateUserPayload

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
		}

		user, err := svc.Create(c.Context(), payload.Name, payload.Email, payload.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(user)
	}
}
