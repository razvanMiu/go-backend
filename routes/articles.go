package routes

import (
	"sirius/controllers"
)

func initArticles() {
	articles := controllers.Articles{}

	api.Get("/articles", articles.GetArticles)
	api.Get("/articles/:id", articles.GetArticle)
	api.Post("/articles/add", articles.Add)
	api.Delete("/articles", articles.DeleteAll)
	api.Delete("/articles/:id", articles.Delete)

}
