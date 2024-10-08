package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/services"
)

func SignUpEmailPassword(c *gin.Context) {
	var signUpDto dto.SignUpEmailPassword

	if err := c.ShouldBindJSON(&signUpDto); err != nil {
		handlers.HandleErrorResponse(c, "invalid request body", err, http.StatusBadRequest)
		return
	}

	user, err := services.SignUpEmailPassword(&signUpDto)

	if err != nil {
		handlers.HandleErrorResponse(c, "failed to create user", err, http.StatusInternalServerError)
		return
	}

	handlers.HandleSuccessWithDataResponse(c, "user signed up", user, http.StatusOK)
}

func SignInEmailPassword(c *gin.Context) {
	var signInDto dto.SignInEmailPassword

	if err := c.ShouldBindJSON(&signInDto); err != nil {
		handlers.HandleErrorResponse(c, "invalid request body", err, http.StatusBadRequest)
		return
	}

	user, err := services.SignInEmailPassword(&signInDto)

	if err != nil {
		handlers.HandleErrorResponse(c, "failed to sign in user", err, http.StatusInternalServerError)
		return
	}

	handlers.HandleSuccessWithDataResponse(c, "user signed in", user, http.StatusOK)
}
