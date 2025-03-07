package handlers

import "github.com/gin-gonic/gin"

/*
{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "age": 25,
  "password": "your_secure_password"
}
*/

func Signup(c *gin.Context) {
	// if you need traceId, take it out from the context

	// c.ShouldBindJSON(&newUser)
	// call the models create user , for now skip validation

	// send the success response
}
