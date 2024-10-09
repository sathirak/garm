package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/handlers"
	"github.com/sathirak/garm/internal/jwt"
	"github.com/sathirak/garm/repository"
)

func Authenticate(c *gin.Context) {

	token, err := jwt.Get(c)

	if err != nil {
		handlers.HandleErrorWithErrorResponse(c, "failed to validate", http.StatusUnauthorized, err)
		return
	}

	if token.ExpiredAt.Before(time.Now()) {
		jwt.Delete(c)
		handlers.HandleErrorWithErrorResponse(c, "token expired", http.StatusUnauthorized, err)
		return
	}

	user, err := repository.GetUserMeta(token.ID)

	if err != nil {
		handlers.HandleErrorWithErrorResponse(c, "failed to get user", http.StatusUnauthorized, err)
		return
	}

	if time.Until(token.ExpiredAt) < (time.Hour * 24 * 15) {

		err = jwt.Set(c, token.ID)

		if err != nil {
			handlers.HandleErrorWithErrorResponse(c, "failed to refresh", http.StatusUnauthorized, err)
			return
		}
	}

	handlers.HandleSuccessWithDataResponse(c, "user validated", user, http.StatusOK)
}
