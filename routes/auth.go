package routes

import (
	"sirius/controllers"

	"github.com/gofiber/contrib/websocket"
)

func initAuth() {
	auth := controllers.Auth{}
	router := api.Group("/auth")

	router.Get("/", auth.GetUser)
	router.Get("/ws", websocket.New(auth.WS))
	router.Post("/signup", auth.Signup)
}
