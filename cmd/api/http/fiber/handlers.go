package fiber

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cashflow"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cmd/api/http/fiber/handler"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cmd/api/http/fiber/middleware"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/healthcheck"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/logger"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payments"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/products"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user"
)

func Handlers(
	appName,
	appVersion string,
	logger logger.Logger,
	healthSvc healthcheck.UseCase,
	authSvc auth.UseCase,
	userSvc user.UseCase,
	productSvc products.UseCase,
	salesSvc sales.UseCase,
	paymentSvc payments.UseCase,
	cashflowSvc cashflow.UseCase,
) http.Handler {
	app := fiber.New(fiber.Config{
		AppName:               appName,
		DisableStartupMessage: true,
	})

	app.Use(
		recover.New(),
		cors.New(),
		requestid.New(),
		middleware.NewOtel(appName),
		middleware.NewLogger(logger, appName, appVersion),
	)

	authMiddleware := middleware.NewAuth(authSvc)

	handler.MakeHealthCheckRoutes(app, healthSvc)
	handler.MakeSwaggerRoutes(app.Group("/docs"))
	handler.MakeAuhtRoutes(app.Group("/auth"), authSvc)
	handler.MakeUserRoutes(app.Group("/user", authMiddleware), userSvc)
	handler.MakeSalesRoutes(app.Group("/sales", authMiddleware), salesSvc)
	handler.MakeProductsRoutes(app.Group("/products", authMiddleware), productSvc)
	handler.MakePaymentsRoutes(app.Group("/payments", authMiddleware), paymentSvc)
	handler.MakeCashFlowHandler(app.Group("/cashflow", authMiddleware), cashflowSvc)

	return adaptor.FiberApp(app)
}
