package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hotelbear/garm/handlers"
	"github.com/hotelbear/garm/internal/errx"
	"github.com/hotelbear/garm/models/dto"
	"github.com/hotelbear/garm/services"
)

func CheckPasswordUser(c *gin.Context) {
	var passwordCheckDTO dto.PasswordCheck

	if err := c.ShouldBindJSON(&passwordCheckDTO); err != nil {
		handlers.Errorx(c, errx.NewError(err, errx.ErrUnprocessableContent))
		return
	}

	passwordCheck := services.PasswordCheck(passwordCheckDTO.Password)

	handlers.SuccessWithDataResponse(c, passwordCheck)
}
