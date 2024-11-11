package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/internal/config"
	"github.com/sathirak/garm/internal/errx"
	"github.com/sathirak/garm/internal/jwt"
	"github.com/sathirak/garm/models"
)

func Authenticate(c *gin.Context, userId *models.UserId) (err errx.Errx) {

	token, err := jwt.Get(c)

	if !err.IsNil() {
		return err
	}

	if token.ExpiredAt.Before(time.Now()) {
		jwt.Delete(c)
		return errx.NewError(err, errx.ErrInvalidToken)
	}

	if time.Until(token.ExpiredAt) < (config.Get().App.JWTExpTime - time.Hour*24) {

		err = jwt.Set(c, token.ID)

		if !err.IsNil() {
			return errx.NewError(err, errx.ErrInvalidToken)
		}
	}
	userId.ID = token.ID
	return errx.Nil()
}
