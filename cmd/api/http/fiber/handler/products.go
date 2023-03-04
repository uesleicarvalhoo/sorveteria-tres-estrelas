package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/product"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/trace"
)

func MakeProductsRoutes(r fiber.Router, svc product.UseCase) {
	r.Get("/:id", getProductByID(svc))
	r.Get("/", getAllProducts(svc))
	r.Post("/", createProduct(svc))
	r.Delete("/:id", deleteProductByID(svc))
	r.Post("/:id", updateProductByID(svc))
}

// @Summary     Get Product by ID
// @Description Get product Data
// @Tags        Product
// @Accept      json
// @Produce     json
// @Param       id path string true "the id of product"
// @Success     200 {object} product.Product
// @Failure     422 {object} dto.MessageJSON "when id is invalid"
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /products/{id} [get]
func getProductByID(svc product.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "get-product-by-id")
		defer span.End()

		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		p, err := svc.Get(ctx, id)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.JSON(p)
	}
}

// @Summary     Get all products
// @Description Get all products data
// @Tags        Product
// @Accept      json
// @Produce     json
// @Success     200 {object} []product.Product
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /products/ [get]
func getAllProducts(svc product.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "get-products")
		defer span.End()

		records, err := svc.Index(ctx)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.JSON(records)
	}
}

// @Summary     Create a New Product
// @Description create a new product and return data
// @Tags        Product
// @Accept      json
// @Produce     json
// @Param       payload body dto.CreateProductPayload true "the product data"
// @Success     200 {object} product.Product
// @Failure     422 {object} dto.MessageJSON "when data is invalid"
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /products/ [post]
func createProduct(svc product.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "create-product")
		defer span.End()

		var payload dto.CreateProductPayload

		if err := c.BodyParser(&payload); err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		p, err := svc.Store(ctx, payload.Name, payload.PriceVarejo, payload.PriceAtacado, payload.AtacadoAmount)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(p)
	}
}

// @Summary     Delete Product by ID
// @Description Delete product
// @Tags        Product
// @Accept      json
// @Produce     json
// @Param       id path string true "the id of product"
// @Success     202
// @Failure     500 {object} dto.MessageJSON "when an error occurs"
// @Router      /products/{id} [delete]
func deleteProductByID(svc product.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "delete-product-by-id")
		defer span.End()

		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: "invalid product id"})
		}

		err = svc.Delete(ctx, id)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.SendStatus(fiber.StatusAccepted)
	}
}

// Summary     Update product by ID
// Description Update product
// Tags        Product
// Accept      json
// Produce     json
// Param       id path string true "the id of product"
// Param       payload body product.UpdatePayload true "the product data"
// Success     200 {object} product.Product
// Failure     422 {object} dto.MessageJSON "when data is invalid"
// Failure     500 {object} dto.MessageJSON "when an error occurs"
// Router      /products/{id} [put]
func updateProductByID(svc product.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := trace.NewSpan(c.UserContext(), "update-product-by-id")
		defer span.End()

		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: "invalid product id"})
		}

		var payload product.UpdatePayload
		if err := c.BodyParser(&payload); err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.MessageJSON{Message: err.Error()})
		}

		p, err := svc.Update(ctx, id, payload)
		if err != nil {
			trace.AddSpanError(span, err)

			return c.Status(fiber.StatusInternalServerError).JSON(dto.MessageJSON{Message: err.Error()})
		}

		return c.Status(http.StatusOK).JSON(p)
	}
}
