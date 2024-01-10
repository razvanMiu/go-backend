package routes

import (
	"sirius/config"
	"sirius/docs"

	"github.com/gofiber/swagger"
)

func initSwagger() {
	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "REST API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.Settings.Host
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Swagger 2.0 routes
	api.Get("/docs/*", swagger.New(swagger.Config{
		DocExpansion: "list",
	}))
}
