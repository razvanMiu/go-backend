package controllers

import (
	"net/http"
	"sirius/config"
	"sirius/models"
	"time"

	"github.com/gin-gonic/gin"
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
func (a Auth) GetUser(c *gin.Context) {
	var user, params models.User

	if err := c.BindUri(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"id": "Invalid user id",
		})
		return
	}

	config.DB.First(&user, "id = ?", params.ID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}

// @Summary Signup
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Router /api/auth/signup [post]
func (a Auth) Signup(c *gin.Context) {
	var user, params models.User

	if err := c.BindUri(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"id": "Invalid user id",
		})
		return
	}

	config.DB.First(&user, "id = ?", params.ID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}

func (a Auth) WS(c *gin.Context) {
	conn, err := config.WSUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	go func() {
		conn.ReadMessage()
	}()
	i := 0
	for {
		i++
		err := conn.WriteJSON(gin.H{
			"i": i,
		})
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
