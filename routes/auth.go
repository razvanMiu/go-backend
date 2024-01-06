package routes

import (
	"sirius/controllers"

	"github.com/gin-gonic/gin"
)

func initAuth(router *gin.Engine) {
	auth := controllers.Auth{}
	router.GET("/api/auth", auth.GetUser)
	routes := router.Group("/api/auth")
	{
		routes.POST("/signup", auth.Signup)
		// routes.POST("/login", auth.Login)
		// routes.PATCH("/patch", auth.Patch)
		routes.GET("/ws", auth.WS)
	}
}
