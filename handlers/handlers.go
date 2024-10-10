package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/pkg/logger"
)

func HandleHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "garm",
	})
}

func HandleSuccessWithDataResponse(c *gin.Context, message string, data interface{}, statusCode int) {
	response := models.Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, response)
}

func HandleSuccessResponse(c *gin.Context, message string, statusCode int) {
	response := models.Response{
		Status:  "success",
		Message: message,
	}
	c.JSON(statusCode, response)
}

func HandleErrorResponse(c *gin.Context, message string, statusCode int) {
	response := models.Response{
		Status:  "error",
		Message: message,
	}
	c.JSON(statusCode, response)
}

func HandleErrorWithErrorResponse(c *gin.Context, message string, statusCode int, err error) {
	logger.Get().Errorw("onprocess", "package", "handler", "error", err.Error())
	response := models.Response{
		Status:  "error",
		Message: message,
	}
	c.JSON(statusCode, response)
}

type UnauthorizedError struct{}

func (e *UnauthorizedError) Error() string {
	return "Unauthorized"
}

func NewUnauthorizedError() *UnauthorizedError {
	return &UnauthorizedError{}
}

func HandleUnauthorisedResponse(c *gin.Context) {
	response := models.Response{
		Status:  "error",
		Message: "Unauthorised",
		Error:   NewUnauthorizedError().Error(),
	}

	c.JSON(http.StatusUnauthorized, response)
}
