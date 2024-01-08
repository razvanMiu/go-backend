package routes

import (
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/blogs", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"results": []gin.H{{"id": "1"}, {"id": "2"}},
		})
	})

	router.GET("/rebuild", func(c *gin.Context) {
		cmd := exec.Command("npm", "run", "build")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = "frontend"
		cmd.Run()

		c.JSON(200, gin.H{
			"message": "rebuild done!",
		})
	})

	initAuth(router)
	initSwagger(router)
	initClient(router)
}
