package fiber

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/http/fiber/handler"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/http/fiber/middleware"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/products"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
)

func Handlers(
	appName,
	appVersion string,
	authSvc auth.UseCase,
	userSvc users.UseCase,
	productSvc products.UseCase,
	salesSvc sales.UseCase,
	balanceSvc balances.UseCase,
) http.Handler {
	app := fiber.New(fiber.Config{
		AppName:               appName,
		DisableStartupMessage: true,
	})

	app.Use(
		recover.New(),
		cors.New(),
		requestid.New(),
	)

	authMiddleware := middleware.NewAuth(authSvc)

	handler.MakeHealthCheckRoutes(app)
	handler.MakeSwaggerRoutes(app.Group("/docs"))
	handler.MakeAuhtRoutes(app.Group("/auth"), authSvc)
	handler.MakeUserRoutes(app.Group("/users", authMiddleware), userSvc)
	handler.MakeBalanceRouter(app.Group("/balances", authMiddleware), balanceSvc)
	handler.MakeProductsRoutes(app.Group("/products", authMiddleware), productSvc)
	handler.MakeSalesRoutes(app.Group("/sales", authMiddleware), salesSvc, balanceSvc)

	return adaptor.FiberApp(app)
}
