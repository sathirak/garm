package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hotelbear/garm/handlers"
)

func Healthz(c *gin.Context) {
	handlers.SuccessWithDataResponse(c, gin.H{"version": "01"})

}
