package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/internal/errx"
)

func Set(c *gin.Context, userId string) errx.Errx {
	bearerToken, err := Generate(userId)
	if !err.IsNil() {
		return err
	}
	c.Header("Authorization", bearerToken)
	return errx.Nil()
}

func Get(c *gin.Context) (*JWT, errx.Errx) {
	bearerToken := c.GetHeader("Authorization")

	jwtData, err := Parse(bearerToken)

	if !err.IsNil() {
		return nil, err
	}

	userJwt := JWT{
		ID:        jwtData.ID,
		ExpiredAt: jwtData.ExpiredAt,
	}

	return &userJwt, errx.Nil()
}

func Delete(c *gin.Context) {
	c.Request.Header.Del("Authorization")
}
