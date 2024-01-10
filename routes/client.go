package routes

import (
	"log"
	"net/http"
	"path"
	"sirius/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

type Client struct {
	http.Dir
}

// Open implements FileSystem using os.Open, opening files for reading rooted
// and relative to the directory d.
func (client Client) Open(name string) (http.File, error) {
	ext := path.Ext(name)
	if name == "/" {
		name = "/index"
	}
	if ext == "" {
		name += ".html"
	}
	f, err := client.Dir.Open(name)
	if err != nil {
		log.Println(err)
	}
	return f, err
}

func initClient() {
	config.App.Static("/", "frontend/public", fiber.Static{
		CacheDuration: -1,
		MaxAge:        0,
	})
	config.App.Static("/_next/static", "frontend/.next/static", fiber.Static{
		CacheDuration: -1,
		MaxAge:        0,
	})
	config.App.Use(
		// func(c *fiber.Ctx) error {
		// 	path := c.Path()
		// 	if StalePaths[path] {
		// 		log.Println("Stale paths", path)
		// 		_, err := http.Get("http://localhost:3000" + path)
		// 		if err != nil {
		// 			log.Println(err)
		// 		}
		// 		delete(StalePaths, path)
		// 	}
		// 	return c.Next()
		// },
		func(c *fiber.Ctx) error {
			log.Println("Client Path = ", c.Path(), c.Get("Rsc") != "")
			if c.Get("Rsc") != "" {
				path := c.Path()
				if path == "/" {
					path = "/index"
				}
				path += ".rsc"
				c.Set(fiber.HeaderContentType, "text/x-component")
				c.Set(fiber.HeaderCacheControl, "no-cache, no-store, must-revalidate")
				c.Set(fiber.HeaderConnection, "keep-alive")
				c.Set(fiber.HeaderVary, "RSC, Next-Router-State-Tree, Next-Router-Prefetch, Next-Url, Accept-Encoding")
				c.Set(fiber.HeaderTransferEncoding, "chunked")
				c.Set(fiber.HeaderKeepAlive, "timeout=5")

				return c.SendFile("frontend/.next/server/app" + path)
			}
			return c.Next()
		},
		filesystem.New(filesystem.Config{
			Root:         Client{http.Dir("frontend/.next/server/app")},
			Browse:       true,
			Index:        "/index.html",
			NotFoundFile: "/_not-found.html",
		}))
}
