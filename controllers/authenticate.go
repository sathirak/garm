package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/services"
)

func Authenticate(c *gin.Context) {

	err := services.Authenticate(c)

	if !err.IsNil() {
		handlers.Errorx(c, err, http.StatusUnauthorized)
		return
	}

	handlers.SuccessResponse(c)
}
