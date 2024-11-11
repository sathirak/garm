package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/internal/errx"
	"github.com/sathirak/garm/internal/jwt"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/services"
)

func SignUpPartner(c *gin.Context) {
	var signUpDTO dto.SignUpPartner

	if err := c.ShouldBindJSON(&signUpDTO); err != nil {
		handlers.Errorx(c, errx.NewError(err, errx.ErrUnprocessableContent))
		return
	}

	user, err := services.SignUpPassword(&signUpDTO)

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

func SignInPartner(c *gin.Context) {
	var signInDTO dto.SignInPartner

	if err := c.ShouldBindJSON(&signInDTO); err != nil {
		handlers.Errorx(c, errx.NewError(err, errx.ErrUnprocessableContent))
		return
	}

	user, err := services.SignInPartner(&signInDTO)

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

func ResetPasswordPartner(c *gin.Context) {
	var resetPasswordDTO dto.ResetPasswordPartner

	if err := c.ShouldBindJSON(&resetPasswordDTO); err != nil {
		handlers.Errorx(c, errx.NewError(err, errx.ErrUnprocessableContent))
		return
	}

	if err := services.ResetPasswordPartner(&resetPasswordDTO, c); !err.IsNil() {
		handlers.Errorx(c, err)
		return
	}

	handlers.SuccessResponse(c)
}
