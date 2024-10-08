package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/pkg/logger"
	"github.com/sathirak/garm/services"
)

func SignUpEmailPassword(c *gin.Context) {
	var signUpDto dto.SignUpEmailPassword
	log := logger.Get();

	if err := c.ShouldBindJSON(&signUpDto); err != nil {
		log.Error(err)
		handlers.HandleErrorResponse(c, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := services.SignUpEmailPassword(&signUpDto)

	if err != nil {
		log.Error(err)
		handlers.HandleErrorResponse(c, "failed to create user", http.StatusInternalServerError)
		return
	}

	handlers.HandleSuccessWithDataResponse(c, "user signed up", user, http.StatusOK)
}

func SignInEmailPassword(c *gin.Context) {
	var signInDto dto.SignInEmailPassword
	log := logger.Get();

	if err := c.ShouldBindJSON(&signInDto); err != nil {
		log.Error(err)
		handlers.HandleErrorResponse(c, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := services.SignInEmailPassword(&signInDto)

	if err != nil {
		log.Error(err)
		handlers.HandleErrorResponse(c, "failed to sign in user", http.StatusInternalServerError)
		return
	}

	if user == nil {
		handlers.HandleSuccessResponse(c, "invalid email or password", http.StatusUnauthorized)
		return
	}

	handlers.HandleSuccessWithDataResponse(c, "user signed in", user, http.StatusOK)
}
