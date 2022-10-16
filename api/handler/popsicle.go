package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/api/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/popsicle"
)

func MakePopsicleRoutes(r fiber.Router, svc popsicle.UseCase) {
	r.Get("/:id", getPopsicleByID(svc))
	r.Get("/", getAllPopsicles(svc))
	r.Post("/", createPopsicle(svc))
	r.Delete("/:id", deletePopsicleByID(svc))
}

// @Summary		Get Popsicle by ID
// @Description	Get popsicle Data
// @Tags		Popsicle
// @Produce		json
// @Param		id			path		string				true	"the id of popsicle"
// @Success		200	{object} entity.Popsicle
// @Failure		422	{object} dto.MessageJSON "when id is invalid"
// @Failure		500	{object} dto.MessageJSON "when an error occurs"
// @Router		/popsicles/{id} [get].
func getPopsicleByID(svc popsicle.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		pop, err := svc.Get(c.Context(), id)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.JSON(pop)
	}
}

// @Summary		Get all popsicles
// @Description	Get all popsicles data
// @Tags		Popsicle
// @Produce		json
// @Success		200	{object} []entity.Popsicle
// @Failure		500	{object} dto.MessageJSON "when an error occurs"
// @Router		/popsicles/ [get].
func getAllPopsicles(svc popsicle.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		records, err := svc.Index(c.Context())
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.JSON(records)
	}
}

// @Summary		Create a New Popsicle
// @Description	create a new popsicle and return data
// @Tags		Popsicle
// @Accept		json
// @Produce		json
// @Param		payload			body		dto.CreatePopsiclePayload				true	"the popsicle data"
// @Success		200	{object} entity.Popsicle
// @Failure		422	{object} dto.MessageJSON "when data is invalid"
// @Failure		500	{object} dto.MessageJSON "when an error occurs"
// @Router		/popsicles/ [post].
func createPopsicle(svc popsicle.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload dto.CreatePopsiclePayload

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		pop, err := svc.Store(c.Context(), payload.Flavor, payload.Price)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.Status(http.StatusCreated).JSON(pop)
	}
}

// @Summary		Delete Popsicle by ID
// @Description	Delete popsicle
// @Tags		Popsicle
// @Produce		json
// @Param		id			path		string				true	"the id of popsicle"
// @Success		202
// @Failure		500	{object} dto.MessageJSON "when an error occurs"
// @Router		/popsicles/{id} [delete].
func deletePopsicleByID(svc popsicle.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		err = svc.Delete(c.Context(), id)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.SendStatus(http.StatusAccepted)
	}
}
