package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/internal/db"
	"github.com/sathirak/garm/models"
)

func NewServiceStatus(service string, err error) models.ServiceStatus {
	status := "ok"
	if err != nil {
		status = "not-ok"
	}
	return models.ServiceStatus{
		Status:    status,
		Service:   service,
		Timestamp: time.Now().Unix(),
	}
}

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
	err := db.Get().Ping()

	dbStatus := NewServiceStatus("database", err)

	details := []models.ServiceStatus{dbStatus}
	handlers.Healthz(c, details)
}
