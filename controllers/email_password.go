package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/internal/jwt"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/repository"
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
		handlers.ErrorWithErrorResponse(c, "invalid request body", http.StatusBadRequest, err)
		return
	}

	if !repository.IsEmailAvailable(signUpDto.Email) {
		handlers.ErrorResponse(c, "email already in use", http.StatusBadRequest)
		return
	}

	user, err := services.SignUpEmailPassword(&signUpDto)

	if err != nil {
		handlers.ErrorWithErrorResponse(c, "failed to create user", http.StatusBadRequest, err)
		return
	}

	if err = jwt.Set(c, user.ID); err != nil {
		handlers.ErrorWithErrorResponse(c, "failed to set auth headers", http.StatusInternalServerError, err)
		return
	}

	handlers.SuccessWithDataResponse(c, user, http.StatusOK)
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
		handlers.ErrorResponse(c, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := services.SignInEmailPassword(&signInDto)

	if err != nil {
		handlers.ErrorWithErrorResponse(c, "failed to sign in user", http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		handlers.ErrorResponse(c, "invalid email or password", http.StatusUnauthorized)
		return
	}

	if err = jwt.Set(c, user.ID); err != nil {
		handlers.ErrorWithErrorResponse(c, "failed to set auth headers", http.StatusInternalServerError, err)
		return
	}

	handlers.SuccessWithDataResponse(c, user, http.StatusOK)
}
