package main

import (
	"log"
	"sirius/config"
	"sirius/routes"
)

func main() {
	config.InitConfig()
	routes.InitRoutes()

	log.Fatal(config.App.Listen(":8080"))
}
