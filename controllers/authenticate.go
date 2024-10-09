package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/internal/jwt"
	"github.com/sathirak/garm/pkg/logger"
)

func Authenticate(c *gin.Context) {
	
	log := logger.Get()

	user, err := jwt.Get(c)

	if err != nil {
		log.Error(err)
		handlers.HandleErrorResponse(c, "failed to validate", http.StatusUnauthorized)
		return
	}

	if user.ExpiredAt.Before(time.Now()) {
		jwt.Delete(c)
		handlers.HandleErrorResponse(c, "token expired", http.StatusUnauthorized)
		return
	}

	if time.Until(user.ExpiredAt) < (time.Hour * 24 * 15) {

		err = jwt.Set(c, user.ID)

		if err != nil {
			log.Error(err)
			handlers.HandleErrorResponse(c, "failed to refresh", http.StatusUnauthorized)
			return
		}
	}

	handlers.HandleSuccessWithDataResponse(c, "user validated", user.User, http.StatusOK)
}
