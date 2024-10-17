package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/models"
)

func SuccessWithDataResponse(c *gin.Context, data interface{}) {
	response := models.Response{
		Status: "ok",
		Data:   data,
	}
	c.JSON(200, response)
}

func SuccessResponse(c *gin.Context) {
	response := models.Response{
		Status: "ok",
	}
	c.JSON(200, response)
}

func ErrorResponse(c *gin.Context, message string, statusCode int) {
	response := models.Response{
		Status:  "error",
		Message: message,
	}
	c.JSON(statusCode, response)
}

func ErrorWithErrorResponse(c *gin.Context, message string, statusCode int, err error) {
	response := models.Response{
		Status:  "error",
		Message: message,
	}
	c.JSON(statusCode, response)
}
