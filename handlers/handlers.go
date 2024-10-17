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

func Errorx(c *gin.Context, err errx.Errx) {

	if svcErr := err.GetSvcError(); svcErr != nil {
		logger.Get().Error(svcErr)
	}

	apiErr := err.GetApiError()
	var genericErr string

	if apiErr != nil {
		genericErr = apiErr.Error()
	} else {
		genericErr = errx.ErrInternalServerErr.Err.Error()
	}

	c.JSON(err.ApiError.StatusCode, models.Response{
		Status:  "error",
		Message: genericErr,
	})
}
