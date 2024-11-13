package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hotelbear/garm/handlers"
	"github.com/hotelbear/garm/internal/errx"
	"github.com/hotelbear/garm/internal/jwt"
	"github.com/hotelbear/garm/models/dto"
	"github.com/hotelbear/garm/services"
)

func SignUpUser(c *gin.Context) {
	var signUpDTO dto.SignUpUser

	if err := c.ShouldBindJSON(&signUpDTO); err != nil {
		handlers.Errorx(c, errx.NewError(err, errx.ErrUnprocessableContent))
		return
	}

	user, err := services.SignUpUser(&signUpDTO)

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

func SignInUser(c *gin.Context) {
	var signInDTO dto.SignInUser

	if err := c.ShouldBindJSON(&signInDTO); err != nil {
		handlers.Errorx(c, errx.NewError(err, errx.ErrUnprocessableContent))
		return
	}

	user, err := services.SignInUser(&signInDTO)

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

func ResetPasswordUser(c *gin.Context) {
	var resetPasswordDTO dto.ResetPasswordUser

	if err := c.ShouldBindJSON(&resetPasswordDTO); err != nil {
		handlers.Errorx(c, errx.NewError(err, errx.ErrUnprocessableContent))
		return
	}

	if err := services.ResetPasswordUser(&resetPasswordDTO, c); !err.IsNil() {
		handlers.Errorx(c, err)
		return
	}

	handlers.SuccessResponse(c)
}
