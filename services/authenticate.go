package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hotelbear/garm/internal/config"
	"github.com/hotelbear/garm/internal/errx"
	"github.com/hotelbear/garm/internal/jwt"
	"github.com/hotelbear/garm/models"
	"github.com/hotelbear/garm/repository"
	"github.com/mitchellh/mapstructure"
)

func Authenticate(c *gin.Context) (*models.UserAuthenticateRes, errx.Errx) {
	var userAuthenticate models.UserAuthenticateRes

	// Get refresh from query string instead of param
	isVerbose := c.Query("verbose") == "true"
	token, err := jwt.Get(c)

	if !err.IsNil() {
		return nil, err
	}

	if token.ExpiredAt.Before(time.Now()) {
		jwt.Delete(c)
		return nil, errx.NewError(err, errx.ErrInvalidToken)
	}

	if time.Until(token.ExpiredAt) < (config.Get().App.JWTExpTime - time.Hour*24) {

		err = jwt.Set(c, token.ID)

		if !err.IsNil() {
			return nil, err
		}
	}

	if isVerbose {
		user, err := repository.GetUserByID(token.ID)

		if !err.IsNil() {
			return nil, err
		}

		// Decode user directly without taking its address
		if err := mapstructure.Decode(user, &userAuthenticate); err != nil {
			return nil, errx.NewError(err, errx.ErrInternalServer)
		}
		return &userAuthenticate, errx.Nil()
	}

	userAuthenticate.ID = token.ID
	return &userAuthenticate, errx.Nil()
}
