package main

import (
	"crud/api/auth"
	"crud/api/users"
	"crud/configurations"
	_ "crud/docs"
	"crud/handlers"
	"crud/middlewares"
	"crud/providers"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	"os"
)

// @title Curd
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	configurations.Config()
	providers.InitDbPostgres()
	app := fiber.New(
		fiber.Config{
			ErrorHandler: handlers.ErrorHandler,
		},
	)

	middlewares.Config(app)
	app.Use(recover2.New())
	app.Use(requestid.New())
	app.Get("/docs/*", swagger.Handler) // default

	app.Get("/docs/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))

	api := app.Group("/api")
	users.NewController(api)
	auth.NewController(api)
	log.Fatal(app.Listen(os.Getenv("PORT")))
}
