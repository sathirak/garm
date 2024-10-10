package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/pkg/logger"
)

func Healthz(c *gin.Context, details []models.ServiceStatus) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"details": details,
	})
}

func SuccessWithDataResponse(c *gin.Context, data interface{}, statusCode int) {
	response := models.Response{
		Status: "success",
		Data:   data,
	}
	c.JSON(statusCode, response)
}

func SuccessResponse(c *gin.Context, statusCode int) {
	response := models.Response{
		Status: "success",
	}
	c.JSON(statusCode, response)
}

func ErrorResponse(c *gin.Context, message string, statusCode int) {
	response := models.Response{
		Status:  "error",
		Message: message,
	}
	c.JSON(statusCode, response)
}

func ErrorWithErrorResponse(c *gin.Context, message string, statusCode int, err error) {
	logger.Get().Errorw("onprocess", "package", "handler", "error", err.Error())
	response := models.Response{
		Status:  "error",
		Message: message,
	}
	c.JSON(statusCode, response)
}

type UnauthorizedError struct{}

// func (e *UnauthorizedError) Unauthorized() string {
// 	return "unauthorized"
// }

func (e *UnauthorizedError) Unauthenticated() string {
	return "unauthorized"
}

func NewAuthorizationError() *UnauthorizedError {
	return &UnauthorizedError{}
}

func UnauthenticatedResponse(c *gin.Context) {
	response := models.Response{
		Status:  "error",
		Message: "unauthorised",
		Error:   NewAuthorizationError().Unauthenticated(),
	}

	c.JSON(http.StatusUnauthorized, response)
}
