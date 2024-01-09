package controllers

import (
	"sirius/models"

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

// func (a Auth) WS(c *fiber.Ctx) error {
// 	if websocket.IsWebSocketUpgrade(c) {
// 		c.Locals("allowed", true)
// 		return c.Next()
// 	}
// 	return fiber.ErrUpgradeRequired
// 	// conn, err := config.WSUpgrader.Upgrade(c.Writer, c.Request, nil)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	// defer conn.Close()
// 	// go func() {
// 	// 	conn.ReadMessage()
// 	// }()
// 	// i := 0
// 	// for {
// 	// 	i++
// 	// 	err := conn.WriteJSON(gin.H{
// 	// 		"i": i,
// 	// 	})
// 	// 	if err != nil {
// 	// 		return
// 	// 	}
// 	// 	time.Sleep(time.Second)
// 	// }
// }
