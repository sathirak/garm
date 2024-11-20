package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hotelbear/garm/handlers"
	"github.com/hotelbear/garm/internal/errx"
	"github.com/hotelbear/garm/models"
	"github.com/hotelbear/garm/services"
)

func CheckPasswordUser(c *gin.Context) {
	var passwordCheckDTO models.PasswordCheckReq

	if err := c.ShouldBindJSON(&passwordCheckDTO); err != nil {
		handlers.Errorx(c, errx.NewError(err, errx.ErrUnprocessableContent))
		return
	}

	passwordCheck := services.CheckPassword(passwordCheckDTO.Password)

	handlers.SuccessWithDataResponse(c, passwordCheck)
}
