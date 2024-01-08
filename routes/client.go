package routes

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type WWW struct {
	d http.Dir
}

func (d WWW) Open(name string) (http.File, error) {
	// Clean path
	path := filepath.Clean(name)
	// Check if file exists
	f, err := d.d.Open(path)
	if os.IsNotExist(err) {
		// Not found, try with .html
		if f, err := d.d.Open(path + ".html"); err == nil {
			return f, nil
		}
		if f, err := d.d.Open("/404.html"); err == nil {
			return f, nil
		}
	}
	return f, err
}

func initClient(router *gin.Engine) {
	router.Use(gin.WrapH((http.FileServer(WWW{http.Dir("frontend/out")}))))
}
