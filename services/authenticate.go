package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/internal/config"
	"github.com/sathirak/garm/internal/errx"
	"github.com/sathirak/garm/internal/jwt"
)

func Authenticate(c *gin.Context) errx.Errx {

	token, err := jwt.Get(c)

	if !err.IsNil() {
		return errx.NewError(err, errx.ErrInvalidToken)
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
	return errx.Nil()
}
