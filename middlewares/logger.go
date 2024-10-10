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

		log.Infow(">>", "method", method, "path", path, "ip", clientIP, "userAgent", userAgent)
		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		statusCode := c.Writer.Status()

		if err := c.Errors.Last(); err != nil {
			errorString := c.Errors.Last().Error()
			log.Errorw(errorString, "method", method, "path", path, "ip", clientIP, "userAgent", userAgent, "latency", latency, "status", statusCode)
		}

		log.Infow("<<", "method", method, "path", path, "ip", clientIP, "userAgent", userAgent, "latency", latency, "status", statusCode)
	}
}
