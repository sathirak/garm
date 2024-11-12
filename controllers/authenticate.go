package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/services"
)

func AuthenticateUser(c *gin.Context) {

	userAuthenticate, err := services.Authenticate(c)

	if !err.IsNil() {
		handlers.Errorx(c, err)
		return
	}

	handlers.SuccessWithDataResponse(c, userAuthenticate)
}
