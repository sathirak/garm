package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/hotelbear/garm/internal/config"
	"github.com/hotelbear/garm/internal/db"
	"github.com/hotelbear/garm/internal/jwt"
	"github.com/hotelbear/garm/internal/logger"
	"github.com/hotelbear/garm/middlewares"
	"github.com/hotelbear/garm/routes"
)

func main() {
	logger.Initialize()
	config.Initialize()
	db.Initialize()
	jwt.Initialize()

	log := logger.Get()
	cfg := config.Get()

	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(requestid.New())
	r.Use(middlewares.Logger())
	routes.SetupRoutes(r)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		db.Close()
		defer logger.Close()
		log.Infow("shutdown", "package", "main", "status", "ok")
		os.Exit(0)
	}()

	if err := r.Run(":" + cfg.App.Port); err != nil {
		log.Errorw("startup", "package", "main", "error", err.Error())
	}
}
