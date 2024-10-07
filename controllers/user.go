package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/services"
)

func CreateUser(c *gin.Context) {
	var dto dto.User
	var user models.User

	if err := c.ShouldBindJSON(&dto); err != nil {
		handlers.HandleErrorResponse(c, "invalid request body", err, http.StatusBadRequest)
		return
	}

	if err := services.CreateUser(&user, dto); err != nil {
		handlers.HandleErrorResponse(c, "failed to create user", err, http.StatusInternalServerError)
		return
	}

	handlers.HandleSuccessWithDataResponse(c, "user created", user, http.StatusOK)
}
