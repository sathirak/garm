package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/repository"
)

func Set(c *gin.Context, userId string) error {
	bearerToken, err := JWTGen(userId)
	if err != nil {
		return err
	}
	c.Header("Authorization", bearerToken)
	return nil
}

func Get(c *gin.Context) (*models.UserJWT, error) {
	bearerToken := c.GetHeader("Authorization")

	jwtData, err := JWTParse(bearerToken)

	if err != nil {
		return nil, err
	}

	user, err := repository.GetUser(jwtData.ID)

	if err != nil {
		return nil, err
	}

	userJwt := models.UserJWT{
		User:      models.User{
			FirstName: user.FirstName, 
			LastName: user.LastName, 
			Email: user.Email, 
			Locale: user.Locale,
		},
		ID:        jwtData.ID,
		ExpiredAt: jwtData.ExpiredAt,
	}

	return &userJwt, nil
}

func Delete(c *gin.Context) {
	c.Request.Header.Del("Authorization")	
}