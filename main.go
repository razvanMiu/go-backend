package main

import (
	"log"
	"sirius/config"
	"sirius/routes"
)

func main() {
	config.InitConfig()
	routes.InitRoutes()

	log.Fatal(config.App.Listen("127.0.0.1:" + config.Settings.Port))
}
