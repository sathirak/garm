package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/internal/errx"
	"github.com/sathirak/garm/internal/jwt"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/services"
)

func SignUpEmailPassword(c *gin.Context) {
	var signUpDto dto.SignUpEmailPassword

	if err := c.ShouldBindJSON(&signUpDto); err != nil {
		handlers.Errorx(c, errx.NewError(err, errx.ErrUnprocessableContent))
		return
	}

	user, err := services.SignUpEmailPassword(&signUpDto)

	if !err.IsNil() {
		handlers.Errorx(c, err)
		return
	}

	if err := jwt.Set(c, user.ID); !err.IsNil() {
		handlers.Errorx(c, err)
		return
	}

	handlers.SuccessWithDataResponse(c, user)
}

func SignInEmailPassword(c *gin.Context) {
	var signInDto dto.SignInEmailPassword

	if err := c.ShouldBindJSON(&signInDto); err != nil {
		handlers.Errorx(c, errx.NewError(err, errx.ErrUnprocessableContent))
		return
	}

	user, err := services.SignInEmailPassword(&signInDto)

	if !err.IsNil() {
		handlers.Errorx(c, err)
		return
	}

	if err := jwt.Set(c, user.ID); !err.IsNil() {
		handlers.Errorx(c, err)
		return
	}

	handlers.SuccessWithDataResponse(c, user)
}

func ResetEmailPassword(c *gin.Context) {
	var resetPasswordDto dto.ResetEmailCredentials

	userID := c.Param("userID")

	if err := c.ShouldBindJSON(&resetPasswordDto); err != nil {
		handlers.Errorx(c, errx.NewError(err, errx.ErrUnprocessableContent))
		return
	}

	if err := services.ResetEmailPassword(&resetPasswordDto, userID); !err.IsNil() {
		handlers.Errorx(c, err)
		return
	}

	handlers.SuccessResponse(c)
}
