package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/internal/errx"
)

func Set(c *gin.Context, userId string) errx.Errx {
	bearerToken, err := Generate(userId)
	if err != nil {
		return errx.NewError(err, errx.ErrInternalServerErr)
	}
	c.Header("Authorization", bearerToken)
	return errx.NewError(nil, nil)
}

func Get(c *gin.Context) (*JWT, errx.Errx) {
	bearerToken := c.GetHeader("Authorization")

	jwtData, err := Parse(bearerToken)

	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	userJwt := JWT{
		ID:        jwtData.ID,
		ExpiredAt: jwtData.ExpiredAt,
	}

	return &userJwt, errx.NewError(err, errx.ErrInternalServerErr)
}

func Delete(c *gin.Context) {
	c.Request.Header.Del("Authorization")
}
