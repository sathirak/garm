package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/pkg/logger"
)

func Logger() gin.HandlerFunc {
	log := logger.Get()

	return func(c *gin.Context) {
		start := time.Now()
		clientIP := c.ClientIP()
		method := c.Request.Method
		userAgent := c.Request.Header.Get("User-Agent")
		path := c.Request.URL.Path

		log.Infow("incoming request", "method", method, "path", path, "ip", clientIP, "userAgent", userAgent)
		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		statusCode := c.Writer.Status()

		if statusCode >= 400 {
			log.Errorw(c.Errors.String(), "method", method, "path", path, "ip", clientIP, "userAgent", userAgent, "latency", latency, "status", statusCode)
		} else {
			log.Infow("outgoing request", "method", method, "path", path, "ip", clientIP, "userAgent", userAgent, "latency", latency, "status", statusCode)

		}
	}
}
