package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/internal/jwt"
	"github.com/sathirak/garm/pkg/logger"
)

func Validate(c *gin.Context) {
	user, err := jwt.Get(c)

	if err != nil {
		logger.Get().Error(err)
		handlers.HandleErrorResponse(c, "failed to validate", http.StatusInternalServerError)
		return
	}

	handlers.HandleSuccessWithDataResponse(c, "user validated", user.User, http.StatusOK)
}
