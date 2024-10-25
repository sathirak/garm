package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/services"
)

// PasswordCheck godoc
//
//	@Summary		Check password
//	@Description	Checks the password strength
//	@Tags			email-password
//	@Accept			json
//	@Produce		json
//	@Param			passwordCheckDTO	body		dto.PasswordCheck	true	"Password Check Data"
//	@Success		200			{object}
//	@Router			/email-password/password-check [post]
//
//	@Security		ApiKeyAuth
func PasswordCheck(c *gin.Context) {
	var passwordCheckDTO dto.PasswordCheck

  passwordCheck := services.PasswordCheck(passwordCheckDTO.Password);

	handlers.SuccessWithDataResponse(c, passwordCheck)
}

