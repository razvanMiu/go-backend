package controllers

import (
	"log"
	"sirius/models"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	models.User
}

// @Summary Get user
// @Description Get user details
// @Tags Auth
// @Accept json
// @Produce json
// @Router /api/auth [get]
func (a Auth) GetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"id": 1, "username": "test"})
}

// @Summary Signup
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Router /api/auth/signup [post]
func (a Auth) Signup(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"signup": "success"})
}

func (a Auth) WS(c *websocket.Conn) {
	closeCh := make(chan bool)

	go func() {
		if _, _, err := c.ReadMessage(); err != nil {
			log.Println("Reading:", err)
			closeCh <- true
		}
	}()

Loop:
	for {
		select {
		case <-closeCh:
			break Loop

		default:
			if err := c.WriteMessage(websocket.TextMessage, []byte("Hello World!")); err != nil {
				log.Println("Writing:", err)
				return
			}
		}
		time.Sleep(time.Second)
	}

}
