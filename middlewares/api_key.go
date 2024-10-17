package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/internal/config"
)

func ApiKeyAuth() gin.HandlerFunc {

	cfg := config.Get()

	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("X-Api-Token")
		if apiKey == "" || apiKey != cfg.App.ApiToken {
			handlers.ErrorWithErrorResponse(c, "invalid api key", 401, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
