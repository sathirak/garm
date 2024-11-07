package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/services"
)

func PasswordCheck(c *gin.Context) {
	var passwordCheckDTO dto.PasswordCheck

	passwordCheck := services.PasswordCheck(passwordCheckDTO.Password)

	handlers.SuccessWithDataResponse(c, passwordCheck)
}
