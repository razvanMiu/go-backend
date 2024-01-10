package routes

import (
	"sirius/config"
)

func initClient() {
	config.App.Static("/_go/static", "public")
}
