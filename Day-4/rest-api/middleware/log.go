package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

type key string

const TraceIdKey key = "1"

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate a new unique identifier (UUID)
		traceId := uuid.NewString()

		requestStartTime := time.Now()

		ctx := c.Request.Context()

		// Add the trace id in context so it can be used by upcoming processes in this request's lifecycle
		ctx = context.WithValue(ctx, TraceIdKey, traceId)

		// The 'WithContext' method on 'c.Request' creates a new copy of the request ('req'),
		// but with an updated context ('ctx') that contains our trace ID.
		// The original request does not get changed by this; we're simply creating a new version of it ('req').

		c.Request = c.Request.WithContext(ctx)

		slog.Info("started", slog.String("Trace ID", traceId),
			slog.String("Method", c.Request.Method), slog.Any("URL Path", c.Request.URL.Path))

		c.Next()

		slog.Info("completed", slog.String("Trace ID", traceId),
			slog.String("Method", c.Request.Method), slog.Any("URL Path", c.Request.URL.Path),
			slog.Int("Status Code", c.Writer.Status()), slog.Int64("duration μs,",
				time.Since(requestStartTime).Microseconds()))

	}
}
