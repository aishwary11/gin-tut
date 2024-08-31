package middleware

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()
		latency := time.Since(startTime)
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		formattedTime := startTime.Format("2006-01-02 15:04:05")
		log.Printf("[GIN] %v | %3d | %v | %s | %s | Body: %s",
			formattedTime,
			statusCode,
			latency,
			method,
			path,
			string(bodyBytes),
		)
	}
}
