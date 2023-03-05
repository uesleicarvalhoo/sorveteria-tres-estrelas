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
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cmd/api/fiber/handler"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/healthcheck"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/infrastructure/http/middleware"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/logger"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payment"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/product"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user"
)

func Handlers(
	appName,
	appVersion string,
	logger logger.Logger,
	authSvc auth.UseCase,
	healthSvc healthcheck.UseCase,
	userSvc user.UseCase,
	productSvc product.UseCase,
	salesSvc sales.UseCase,
	paymentSvc payment.UseCase,
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
		middleware.NewFiberOtel(appName),
		middleware.NewFiberLogger(logger, appName, appVersion),
	)

	handler.MakeHealthCheckRoutes(app, healthSvc)
	handler.MakeSwaggerRoutes(app.Group("/docs"))
	handler.MakeAuhtRoutes(app.Group("/auth"), authSvc)
	handler.MakeUserRoutes(app.Group("/user"), userSvc)
	handler.MakeSalesRoutes(app.Group("/sales"), salesSvc)
	handler.MakeProductsRoutes(app.Group("/products"), productSvc)
	handler.MakePaymentsRoutes(app.Group("/payments"), paymentSvc)
	handler.MakeCashFlowHandler(app.Group("/cashflow"), cashflowSvc)

	return adaptor.FiberApp(app)
}
