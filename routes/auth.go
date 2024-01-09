package routes

import (
	"sirius/config"
	"sirius/controllers"
)

func initAuth() {
	auth := controllers.Auth{}
	router := config.App.Group("/api/auth")

	router.Get("/", auth.GetUser)
	// router.Get("/ws", auth.WS)
	router.Post("/signup", auth.Signup)
}
