package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/internal/config"
	"github.com/sathirak/garm/internal/jwt"
	"github.com/sathirak/garm/middlewares"
	"github.com/sathirak/garm/routes"

	_ "github.com/sathirak/garm/docs"
	"github.com/sathirak/garm/internal/db"
	"github.com/sathirak/garm/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Garm API Docs
//	@version		1.0
//	@description	OpenAPI Docs for Garm Auth Server

//	@host		localhost:9000
//	@BasePath	/api/v1/auth/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						x-api-token

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

	if cfg.App.Env == "development" {
		r.GET("/api/v1/auth/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Use(requestid.New())
	r.Use(middlewares.Logger())
	r.Use(middlewares.ApiKeyAuth())
	routes.SetupRoutes(r)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		db.Close()
		log.Infow("shutdown", "package", "main", "status", "ok")
		os.Exit(0)
	}()

	if err := r.Run(":" + cfg.App.Port); err != nil {
		log.Errorw("startup", "package", "main", "error", err.Error())
	}
}
