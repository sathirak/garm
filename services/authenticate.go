package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/internal/config"
	"github.com/sathirak/garm/internal/errx"
	"github.com/sathirak/garm/internal/jwt"
	"github.com/sathirak/garm/models"
)

func Authenticate(c *gin.Context) (*models.UserAuthenticate, errx.Errx) {
	var userAuthenticate models.UserAuthenticate
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
			return nil, errx.NewError(err, errx.ErrInvalidToken)
		}
	}
	userAuthenticate.ID = token.ID
	return &userAuthenticate, errx.Nil()
}
