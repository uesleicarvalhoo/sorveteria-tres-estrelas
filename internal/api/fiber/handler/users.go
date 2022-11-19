package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
)

func MakeUserRoutes(r fiber.Router, userSvc users.UseCase) {
	r.Get("/me", getMe(userSvc))
	r.Post("/", createUser(userSvc))
}

// @Summary		Get Me
// @Description	Get current user data
// @Tags		User
// @Accept      json
// @Produce     json
// @Success		200	{object} users.User
// @Failure		500	{object} dto.MessageJSON "when an error occurs"
// @Router		/users/me [get]
func getMe(svc users.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctxID := c.Locals("userID")

		id, ok := ctxID.(uuid.UUID)
		if !ok {
			return c.Status(http.StatusInternalServerError).JSON(dto.MessageJSON{Message: "user data not found"})
		}

		u, err := svc.Get(c.Context(), id)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.JSON(u)
	}
}

// @Summary		Create User
// @Description	Create a new user and return user data
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		payload		body	dto.CreateUserPayload true "the user data"
// @Success		201	{object} users.User
// @Failure		422	{object} dto.MessageJSON "when payload is invalid"
// @Failure		500	{object} dto.MessageJSON "when an error occurs"
// @Router		/users [post]
func createUser(svc users.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload dto.CreateUserPayload

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
		}

		id, ok := c.Locals("userID").(uuid.UUID)
		if !ok {
			return c.Status(http.StatusInternalServerError).JSON(dto.MessageJSON{Message: "user data not found"})
		}

		currentUser, err := svc.Get(c.Context(), id)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		permissions := []users.Permission{}

		for _, payloadPerm := range payload.Permissions {
			for i := range currentUser.Permissions {
				if currentUser.Permissions[i] == payloadPerm {
					permissions = append(permissions, payloadPerm)
				}
			}
		}

		user, err := svc.Create(c.Context(), payload.Name, payload.Email, payload.Password, permissions...)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}

		return c.Status(http.StatusCreated).JSON(user)
	}
}
