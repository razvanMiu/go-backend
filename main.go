package main

import (
	"log"
	"path/filepath"
	"runtime"
	"sirius/config"
	"sirius/routes"
)

var Basepath string

func main() {
	_, b, _, _ := runtime.Caller(0)
	Basepath = filepath.Dir(b)

	config.InitConfig()
	routes.InitRoutes()

	log.Fatal(config.App.Listen("0.0.0.0:" + config.Settings.Port))
}
