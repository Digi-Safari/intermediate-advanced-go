package handlers

import (
	"github.com/gin-gonic/gin"
	"rest-api/auth"
	"rest-api/middleware"
)

func API(a *auth.Auth) (*gin.Engine, error) {
	r := gin.New()
	m, err := middleware.NewMid(a)
	if err != nil {
		return nil, err
	}
	r.Use(middleware.Logger(), m.Authenticate())
	r.GET("/check", Check)

	return r, nil
}

func Check(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
