package main

import (
	"os"
	"sirius/config"
	"sirius/routes"

	"github.com/gin-gonic/gin"
)

func InitGinEngine() *gin.Engine {
	gin.SetMode(os.Getenv(gin.EnvGinMode))

	return gin.Default()
}

func main() {
	config.InitConfig()

	ginEngine := InitGinEngine()

	ginEngine.SetTrustedProxies([]string{"localhost:3000"})

	routes.InitRoutes(ginEngine)

	ginEngine.Run()
}
