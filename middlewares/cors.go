package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hotelbear/garm/internal/config"
)

func Cors() gin.HandlerFunc {
	cfg := config.Get()

	origins := []string{
		"https://hotelbear.lk",
	}

	if cfg.App.Env != "production" {
		origins = append(origins, "http://localhost:3000")
		origins = append(origins, "https://stg.hotelbear.lk")
	}

	return cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Forwarded-For"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
