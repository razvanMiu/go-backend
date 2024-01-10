package routes

import (
	"sirius/config"

	"github.com/gofiber/fiber/v2"
)

var api fiber.Router

func InitRoutes() {
	api = config.App.Group("/api")

	initAuth()
	initSwagger()
	initClient()
	initArticles()
}
