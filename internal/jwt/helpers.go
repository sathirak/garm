package jwt

import (
	"github.com/gin-gonic/gin"
)

func Set(c *gin.Context, userId string) error {
	bearerToken, err := JWTGen(userId)
	if err != nil {
		return err
	}
	c.Header("Authorization", bearerToken)
	return nil
}

func Get(c *gin.Context) (*JWT, error) {
	bearerToken := c.GetHeader("Authorization")

	jwtData, err := JWTParse(bearerToken)

	if err != nil {
		return nil, err
	}

	userJwt := JWT{
		ID:        jwtData.ID,
		ExpiredAt: jwtData.ExpiredAt,
	}

	return &userJwt, nil
}

func Delete(c *gin.Context) {
	c.Request.Header.Del("Authorization")
}
