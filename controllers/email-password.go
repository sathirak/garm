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

func SignUpEmailPassword(c *gin.Context) {
	var signUpDto dto.SignUpEmailPassword

	if err := c.ShouldBindJSON(&signUpDto); err != nil {
		handlers.HandleErrorWithErrorResponse(c, "invalid request body", http.StatusBadRequest, err)
		return
	}

	if !repository.CheckUserAvailablityEmail(signUpDto.Email) {
		handlers.HandleSuccessResponse(c, "email already in use", http.StatusConflict)
		return
	}

	user, err := services.SignUpEmailPassword(&signUpDto)

	if err != nil {
		handlers.HandleErrorWithErrorResponse(c, "failed to create user", http.StatusInternalServerError, err)
		return
	}

	err = jwt.Set(c, user.ID)

	if err != nil {
		handlers.HandleErrorWithErrorResponse(c, "failed to set auth headers", http.StatusInternalServerError, err)
		return
	}

	handlers.HandleSuccessWithDataResponse(c, "user signed up", user, http.StatusOK)
}

func SignInEmailPassword(c *gin.Context) {
	var signInDto dto.SignInEmailPassword

	if err := c.ShouldBindJSON(&signInDto); err != nil {
		handlers.HandleSuccessResponse(c, "invalid request body", http.StatusBadRequest)
		return
	}

	if repository.CheckUserAvailablityEmail(signInDto.Email) {
		handlers.HandleErrorResponse(c, "email not found", http.StatusConflict)
		return
	}

	user, err := services.SignInEmailPassword(&signInDto)

	if err != nil {
		handlers.HandleErrorWithErrorResponse(c, "failed to sign in user", http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		handlers.HandleSuccessResponse(c, "invalid email or password", http.StatusUnauthorized)
		return
	}

	err = jwt.Set(c, user.ID)

	if err != nil {
		handlers.HandleErrorWithErrorResponse(c, "failed to set auth headers", http.StatusInternalServerError, err)
		return
	}

	handlers.HandleSuccessWithDataResponse(c, "user signed in", user, http.StatusOK)
}
