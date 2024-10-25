package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/internal/errx"
	"github.com/sathirak/garm/internal/jwt"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/services"
)

// SignUpEmailPassword godoc
//
//	@Summary		Sign up with email and password
//	@Description	Creates a new user account using email and password
//	@Tags			email-password
//	@Accept			json
//	@Produce		json
//	@Param			signUpDto	body		dto.SignUpEmailPassword	true	"Sign Up Data"
//	@Success		200			{object}	models.UserMeta
//	@Router			/email-password/sign-up [post]
//
//	@Security		ApiKeyAuth
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

// SignInEmailPassword godoc
//
//	@Summary		Sign In with email and password
//	@Description	Returns a JWT token for the user
//	@Tags			email-password
//	@Accept			json
//	@Produce		json
//	@Param			signUpDto	body		dto.SignInEmailPassword	true	"Sign In Data"
//	@Success		200			{object}	models.UserMeta
//	@Router			/email-password/sign-in [post]
//
//	@Security		ApiKeyAuth
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

// ResetEmailPassword godoc
//
//	@Summary		Reset password
//	@Description	Resets the password for the user
//	@Tags			email-password
//	@Accept			json
//	@Produce		json
//	@Param			userID			path		string						true	"User ID"
//	@Param			resetPasswordDto	body		dto.ResetEmailCredentials	true	"Reset Password Data"
//	@Success		200			{object}	models.UserMeta
//	@Router			/email-password/reset/{userID} [post]
//
//	@Security		ApiKeyAuth
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
