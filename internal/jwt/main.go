package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/repository"
)

func Set(c *gin.Context, user *models.UserMeta) error {
	token, err := JWTGen(user.ID)
	if err != nil {
		return err
	}
	c.Header("Authorization", token)
	return nil
}

func Get(c *gin.Context) (*models.UserJWT, error) {
	token := c.GetHeader("Authorization")
	jwtData, err := JWTParse(token)

	if err != nil {
		return nil, err
	}

	user, err := repository.GetUser(jwtData.ID)

	if err != nil {
		return nil, err
	}

	userJwt := models.UserJWT{
		User:      models.User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Locale: user.Locale},
		ID:        jwtData.ID,
		ExpiredAt: jwtData.ExpiredAt,
	}

	return &userJwt, nil
}
