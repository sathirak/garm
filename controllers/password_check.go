package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/internal/errx"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/services"
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
