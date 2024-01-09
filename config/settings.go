package config

import (
	"fmt"
	"os"
)

var Settings struct {
	Development bool
	Host        string
	Port        string
}

func initSettings() {
	// Get environment variables
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	// Set development mode
	development := os.Getenv("MODE") == "debug"

	// Set host
	if development {
		host = fmt.Sprintf("localhost:%s", port)
	} else if host == "" {
		host, _ = os.Hostname()
	}

	// Set port
	if port == "" {
		port = "8080"
	}

	// Set Settings
	Settings.Host = host
	Settings.Port = port
	Settings.Development = development
}
