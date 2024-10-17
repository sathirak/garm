package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/internal/config"
	"github.com/sathirak/garm/internal/errx"
)

func ApiKeyAuth() gin.HandlerFunc {

	cfg := config.Get()

	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("X-Api-Token")
		if apiKey == "" || apiKey != cfg.App.ApiToken {
			err := errx.NewError(nil, errx.ErrMissingOrMalformedApiToken)
			handlers.Errorx(c, err, 401)
			c.Abort()
			return
		}
		c.Next()
	}
}
