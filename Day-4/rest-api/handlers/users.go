package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"log/slog"
	"net/http"
	"rest-api/middleware"
	"rest-api/models"
)

type handler struct {
	conn     models.Service
	validate *validator.Validate
}

/*
{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "age": 25,
  "password": "your_secure_password"
}
*/

func (h *handler) Signup(c *gin.Context) {

	// if you need traceId, take it out from the context
	traceId := GetTraceIdOfRequest(c)

	// Check if the size of body is more than 5KB
	if c.Request.ContentLength > 5*1024 {
		slog.Error("request body limit", slog.String("TraceID", traceId), slog.Int64("Size Received", c.Request.ContentLength))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Request body too large. Limit is 5KB"})
		return
	}

	// read the jsonBody, and put it inside the struct
	var newUser models.NewUser
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		slog.Error("json validation error", slog.String("TraceID", traceId),
			slog.String("Error", err.Error()))

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(newUser)
	err = h.validate.Struct(newUser)
	if err != nil {
		slog.Error("json validation error", slog.String("TraceID", traceId),
			slog.String("Error", err.Error()))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//json.NewDecoder(c.Request.Body).Decode(&newUser)
	//
	//conn := models.NewConn() // not the way to Go
	usr, err := h.conn.CreateUser(newUser)
	if err != nil {
		slog.Error("error in creating the user", slog.String("TraceID", traceId),
			slog.String("Error", err.Error()))
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	// call the models create user , for now skip validation

	// send the success response
	// Respond with user data
	c.JSON(http.StatusOK, usr)
}

func GetTraceIdOfRequest(c *gin.Context) string {

	// We get the current request context
	ctx := c.Request.Context()

	// Extract the traceId from the request context
	// We assert the type to string since context.Value returns an interface{}
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)

	// If traceId not present then log the error and return an error message
	// ok is false if the type assertion was not successful
	if !ok {
		slog.Error("trace id not present in the context")
		traceId = "Unknown"
	}
	return traceId
}

func (h *handler) GetUser(c *gin.Context) {
	traceId := GetTraceIdOfRequest(c)

	// Fetch the user_id from the url parameters
	userEmail := c.Query("user_email")

	err := h.validate.Var(userEmail, "required,email")
	if err != nil {
		slog.Error("invalid email", slog.String("TraceID", traceId),
			slog.String("Error", err.Error()))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing email"})
		return
	}

	user, err := h.conn.FetchUser(userEmail)
	if err != nil {
		slog.Error("user not found", slog.String("TraceID", traceId),
			slog.String("Error", err.Error()))

		fetchError := gin.H{"msg": "user not found"}
		// Respond with an StatusNotFound error code and the error message
		c.AbortWithStatusJSON(http.StatusNotFound, fetchError)
		return
	}

	c.JSON(http.StatusOK, user)

}
