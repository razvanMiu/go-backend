package config

import "github.com/gofiber/fiber/v2"

var App *fiber.App

func InitConfig() {
	loadEnv()
	initDB()
	initSettings()
	initApp()
}

func initApp() {
	App = fiber.New()
}
