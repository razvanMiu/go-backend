package routes

import (
	"sirius/config"
	"sirius/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initSwagger(router *gin.Engine) {
	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "REST API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.Settings.Host
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Swagger 2.0 routes
	router.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
