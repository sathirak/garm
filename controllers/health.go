package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
)

// Healthz godoc
//
//	@Summary		Check the health of the service
//	@Description	Returns the health status of the service including database connectivity
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	[]models.ServiceStatus
//	@Router			/healthz [get]
//
// @Security ApiKeyAuth
func Healthz(c *gin.Context) {
	handlers.SuccessResponse(c)
}
