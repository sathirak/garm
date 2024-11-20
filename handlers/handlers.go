package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hotelbear/garm/internal/errx"
	"github.com/hotelbear/garm/internal/logger"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessWithDataResponse(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Status: "ok",
		Data:   data,
	})
}

func SuccessResponse(c *gin.Context) {
	c.JSON(200, Response{
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
		genericErr = errx.ErrInternalServer.Err.Error()
	}

	c.JSON(err.ApiError.StatusCode, Response{
		Status:  "error",
		Message: genericErr,
	})
}
