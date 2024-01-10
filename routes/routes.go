package routes

import (
	"log"
	"net/http"
	"sirius/config"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Article struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

var api fiber.Router

// var StalePaths = make(map[string]bool)

var i = 0
var articles = []Article{}

func revalidate(path string) {
	http.Get("http://localhost:3000/api/revalidate?path=" + path)
	http.Get("http://localhost:3000" + path)

	// StalePaths[path] = true
}

func InitRoutes() {
	api = config.App.Group("/api")

	api.Get("/articles", func(c *fiber.Ctx) error {
		log.Println("Fetch articles...")

		return c.JSON(articles)
	})

	api.Get("/articles/revalidate", func(c *fiber.Ctx) error {
		log.Println("Revalidate articles...")

		i++
		articles = append(articles, Article{Id: strconv.Itoa(i), Title: "Article " + strconv.Itoa(i)})

		go revalidate("/")
		for _, article := range articles {
			go revalidate("/articles/" + article.Id)
		}

		return c.SendStatus(200)
	})

	initAuth()
	initSwagger()
	initClient()
}
