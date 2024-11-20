package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hotelbear/garm/handlers"
	"github.com/hotelbear/garm/internal/config"
)

func Healthz(c *gin.Context) {
	handlers.SuccessWithDataResponse(c, gin.H{"version": config.Get().App.ApiVersion})

}
