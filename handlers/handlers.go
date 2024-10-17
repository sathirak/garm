package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/internal/errx"
	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/pkg/logger"
)

func SuccessWithDataResponse(c *gin.Context, data interface{}) {
	c.JSON(200, models.Response{
		Status: "ok",
		Data:   data,
	})
}

func SuccessResponse(c *gin.Context) {
	c.JSON(200, models.Response{
		Status: "ok",
	})
}

func ErrorResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, models.Response{
		Status:  "error",
		Message: message,
	})
}

func ErrorWithErrorResponse(c *gin.Context, message string, statusCode int, err error) {
	c.JSON(statusCode, models.Response{
		Status:  "error",
		Message: message,
	})
}

func Errorx(c *gin.Context, err errx.Errx, statusCode int) {

  if svcErr := err.GetSvcError(); svcErr != nil {
    logger.Get().Error(svcErr)
  }

  apiErr := err.GetApiError()
  var genericErr string

  if apiErr != nil {
    genericErr = apiErr.Error()
  } else {
    genericErr = errx.ErrInternalServerErr.Error()
  }

  c.JSON(statusCode, models.Response{
    Status:  "error",
    Message: genericErr,
  })
}
