package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/services"
)

func AuthenticateUser(c *gin.Context) {
	var userId models.UserId
	err := services.Authenticate(c, &userId)

	if !err.IsNil() {
		handlers.Errorx(c, err)
		return
	}

	handlers.SuccessWithDataResponse(c, userId)
}
