package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"rest-api/auth"
	"rest-api/middleware"
	"rest-api/models"
)

func API(a *auth.Auth, conn models.Conn) (*gin.Engine, error) {
	r := gin.New()
	m, err := middleware.NewMid(a)
	if err != nil {
		return nil, err
	}
	h := handler{
		conn:     conn,
		validate: validator.New(),
	}
	r.Use(middleware.Logger())
	r.POST("/signup", h.Signup)

	r.Use(m.Authenticate())
	r.GET("/check", Check)
	return r, nil
}

func Check(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
