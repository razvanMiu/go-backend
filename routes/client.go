package routes

import (
	"net/http"
	"os"
	"path/filepath"
	"sirius/config"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

type WWW struct {
	d http.Dir
}

func (d WWW) Open(name string) (http.File, error) {
	// Clean path
	path := filepath.Clean(name)
	// Check if file exists as html
	f, err := d.d.Open(path + ".html")
	if os.IsNotExist(err) {
		// Not found, try to open the original path
		if f, err := d.d.Open(path); err == nil {
			return f, nil
		}
		// TODO: if path is folder, try to open index.html inside
	}
	return f, err
}

func initClient() {
	config.App.Static("/", "frontend/public")
	config.App.Static("/_next/static", "frontend/.next/static")
	config.App.Use(filesystem.New(filesystem.Config{
		Root:         WWW{http.Dir("frontend/.next/server/app")},
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "_not-found.html",
	}))
}
