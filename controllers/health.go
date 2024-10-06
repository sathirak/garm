package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
)

func Healthz(c *gin.Context) {
	handlers.HandleSuccessResponse(c, "Garm is healthy", http.StatusOK)
}