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
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api/fiber/handler"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api/fiber/middleware"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/pkg/logger"
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
	logger logger.Logger,
) http.Handler {
	app := fiber.New(fiber.Config{
		AppName:               appName,
		DisableStartupMessage: true,
	})

	logrusMiddleware := middleware.NewLogrus(logger, appName, appVersion)

	app.Use(
		recover.New(),
		cors.New(),
		requestid.New(),
		logrusMiddleware,
	)

	handler.MakeHealthCheckRoutes(app)
	handler.MakeSwaggerRoutes(app.Group("/docs"))
	handler.MakeAuhtRoutes(app.Group("/auth"), authSvc)
	handler.MakeUserRoutes(app.Group("/users", middleware.NewAuth(authSvc, "users")), userSvc)
	handler.MakeSalesRoutes(
		app.Group("/sales", middleware.NewAuth(authSvc, "sales")), salesSvc, balanceSvc)
	handler.MakeProductsRoutes(app.Group("/products", middleware.NewAuth(authSvc, "products")), productSvc)
	handler.MakeBalanceRouter(app.Group("/balances", middleware.NewAuth(authSvc, "balances")), balanceSvc)

	return adaptor.FiberApp(app)
}
