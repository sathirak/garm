package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
)

func Healthz(c *gin.Context) {
	handlers.HandleHealth(c)
}