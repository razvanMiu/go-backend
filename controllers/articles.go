package controllers

import (
	"encoding/json"
	"log"
	"os"

	"sirius/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Article struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type Articles struct{}

func getArticles() []Article {
	data, err := os.ReadFile("public/articles.json")
	var articles []Article

	if err != nil {
		return articles
	}

	err = json.Unmarshal(data, &articles)

	if err != nil {
		return articles
	}

	return articles
}

func revalidateArticles(articles []Article) {
	go utils.Revalidate("/")
	for _, article := range articles {
		go utils.Revalidate("/articles/" + article.Id)
	}
}

//	@Summary		Get articles
//	@Description	Get list of articles
//	@Tags			Articles
//	@Produce		json
//	@Success		200	{array}	Article
//	@Router			/api/articles [get]
func (a Articles) GetArticles(c *fiber.Ctx) error {
	articles := getArticles()

	return c.JSON(articles)
}

//	@Summary		Get article
//	@Description	Get article by id
//	@Tags			Articles
//	@Produce		json
//	@Param			id	path		string	true	"Id"
//	@Success		200	{object}	Article
//	@Router			/api/articles/{id} [get]
func (a Articles) GetArticle(c *fiber.Ctx) error {
	articles := getArticles()

	id := c.Params("id")

	for _, article := range articles {
		if article.Id == id {
			return c.JSON(article)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

//	@Summary		Add article
//	@Description	Add new article
//	@Tags			Articles
//	@Produce		json
//	@Param			title	formData	string	true	"Title"
//	@Success		200		{array}		Article
//	@Error			400 {object} string
//	@Error			500 {object} string
//	@Router			/api/articles/add [post]
func (a Articles) Add(c *fiber.Ctx) error {
	articles := getArticles()
	var article Article
	if err := c.BodyParser(&article); err != nil {
		log.Println(err)
		// TODO: Must validate
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if article.Title == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	article.Id = uuid.New().String()
	articles = append(articles, article)

	data, _ := json.MarshalIndent(articles, "", "  ")

	err := os.WriteFile("public/articles.json", data, 0o666)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	go revalidateArticles(articles)

	return c.JSON(article)
}

//	@Summary		Delete articles
//	@Description	Delete all articles
//	@Tags			Articles
//	@Produce		json
//	@Success		200	{object}	string
//	@Error			404	{object} string
//	@Error			500	{object} string
//	@Router			/api/articles [delete]
func (a Articles) DeleteAll(c *fiber.Ctx) error {
	articles := getArticles()
	id := c.Params("id")
	update := false
	for i, article := range articles {
		if article.Id == id {
			articles = append(articles[:i], articles[i+1:]...)
			update = true
			go revalidateArticles(articles)
			break
		}
	}
	if update {
		data, _ := json.MarshalIndent(articles, "", "  ")

		err := os.WriteFile("public/articles.json", data, 0o666)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusOK)
	}
	return c.SendStatus(fiber.StatusNotFound)
}

//	@Summary		Delete article
//	@Description	Delete article
//	@Tags			Articles
//	@Produce		json
//	@Param			id	path		string	true	"Id"
//	@Success		200	{object}	string
//	@Success		500	{object}	string
//	@Router			/api/articles/{id} [delete]
func (a Articles) Delete(c *fiber.Ctx) error {
	articles := getArticles()
	data, _ := json.MarshalIndent([]Article{}, "", "  ")

	err := os.WriteFile("public/articles.json", data, 0o666)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	revalidateArticles(articles)

	return c.SendStatus(fiber.StatusOK)
}