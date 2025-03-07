package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"rest-api/auth"
	"strings"
)

func (m *Mid) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		traceId, ok := ctx.Value(TraceIdKey).(string)
		if !ok {
			slog.Error("trace id not present in the context")
			traceId = "unknown"
		}

		authHeader := c.Request.Header.Get("Authorization")

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			// If the header format doesn't match required format, log and send an error
			err := errors.New("expected authorization header format: Bearer <token>")
			slog.Error("invalid authorization header format", slog.String("TraceId", traceId),
				slog.String("Error", err.Error()))

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return // don't forget this
		}
		token := parts[1]

		claims, err := m.a.ValidateToken(token)
		if err != nil {
			slog.Error("invalid token", slog.String("TraceId", traceId),
				slog.String("Error", err.Error()))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
			return
		}

		ctx = context.WithValue(ctx, auth.Key, claims)
		c.Request = c.Request.WithContext(ctx)
		c.Next()

	}
}
