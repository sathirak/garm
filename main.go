package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/internal/config"
	"github.com/sathirak/garm/internal/jwt"
	"github.com/sathirak/garm/middlewares"
	"github.com/sathirak/garm/routes"

	"github.com/sathirak/garm/internal/db"
	"github.com/sathirak/garm/pkg/logger"
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

	r.Use(middlewares.Logger())
	r.Use(middlewares.ApiKeyAuth())
	routes.SetupRoutes(r)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		db.Close()
		log.Infow("shutdown", "package", "garm", "status", "ok")
		os.Exit(0)
	}()

	err := r.Run(":" + cfg.App.Port)
	if err != nil {
		log.Error(err.Error())
	}
}
