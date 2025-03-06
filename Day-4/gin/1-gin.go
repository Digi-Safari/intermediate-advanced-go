package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//r:= gin.Default()
	r := gin.New()
	r.GET("/json", sendJson)

	// Route Parameters
	// name is going to be a parameter
	r.Use(gin.Logger())
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name") // fetching the param
		c.String(200, "Hello, %s! (Gin)", name)
	})

	// Grouping Routes
	v1 := r.Group("/v1")
	{

		v1.GET("/users", func(c *gin.Context) {
			panic("something went wrong in users")
			c.String(200, "Users v1 (Gin)")
		})
		v1.Use(gin.Recovery(), Mid2())
		v1.GET("/posts", func(c *gin.Context) {
			panic("something went wrong in posts")
			c.String(200, "Posts v1 (Gin)")
		})
	}
	r.Run(":8080")
}

func sendJson(c *gin.Context) {
	u := struct {
		Message string `json:"message"`
	}{
		Message: "Hello, World!",
	}

	// convert the struct to json and send the response as well
	// it would set the content type automatically
	c.JSON(http.StatusOK, u)

}

func Mid(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware")
		next(c)
	}
}

func Mid2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
