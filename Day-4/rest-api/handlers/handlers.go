package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"rest-api/auth"
	"rest-api/middleware"
	"rest-api/models"
)

func API(a *auth.Auth, conn models.Service) (*gin.Engine, error) {
	r := gin.New()
	m, err := middleware.NewMid(a)
	if err != nil {
		return nil, err
	}

	// injecting dependencies to handlers package by using a struct and adding required fields
	h := handler{
		conn:     conn,
		validate: validator.New(),
	}

	r.Use(middleware.Logger())
	r.POST("/signup", h.Signup)

	// authenticate would be only applied to check
	r.Use(m.Authenticate())
	r.GET("/check", Check)
	r.GET("/fetch", h.GetUser)
	return r, nil
}

func Check(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
