package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cmd/api/fiber/docs" // Load swagger docs
)

func MakeSwaggerRoutes(r fiber.Router) {
	r.Get("/swagger/*", swagger.HandlerDefault)
}
