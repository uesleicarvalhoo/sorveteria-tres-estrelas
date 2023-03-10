package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/server/http/docs" // Load swagger docs
)

func Swagger(r fiber.Router) {
	r.Get("/swagger/*", swagger.HandlerDefault)
}
