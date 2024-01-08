package routes

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.Engine) {
	initAuth(router)
	initSwagger(router)
	initClient(router)
}
