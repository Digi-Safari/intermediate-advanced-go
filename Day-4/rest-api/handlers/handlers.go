package handlers

import (
	"github.com/gin-gonic/gin"
	"rest-api/middleware"
)

func API() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger())
	r.GET("/check", Check)

	return r
}

func Check(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
