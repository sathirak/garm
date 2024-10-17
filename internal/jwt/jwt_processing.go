package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sathirak/garm/internal/config"
	"github.com/sathirak/garm/internal/errx"
)

type JWT struct {
	ID        string
	ExpiredAt time.Time
}

func Generate(id string) (string, errx.Errx) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(config.Get().App.JWTExpTime).Unix(),
	})

	tokenString, err := token.SignedString(GetKey())
	if err != nil {
		return "", errx.NewError(err, errx.ErrInternalServerErr)
	}

	bearerToken := "Bearer " + tokenString

	return bearerToken, errx.Nil()
}

func Parse(bearerToken string) (jwtData *JWT, err errx.Errx) {

	prefix := "Bearer "
	if !(len(bearerToken) > len(prefix) && bearerToken[:len(prefix)] == prefix) {
		return nil, errx.NewError(nil, errx.ErrUnauthenticated)
	}

	token, newErr := jwt.Parse(bearerToken[len(prefix):], func(token *jwt.Token) (interface{}, error) {
		return GetKey(), nil
	})

	if newErr != nil {
		return nil, errx.NewError(newErr, errx.ErrUnauthenticated)
	}

	if !token.Valid {
		return nil, errx.NewError(nil, errx.ErrUnauthenticated)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errx.NewError(nil, errx.ErrUnauthenticated)
	}

	id := claims["id"].(string)
	exp := claims["exp"].(float64)
	expiredAt := time.Unix(int64(exp), 0)

	jwtData = &JWT{ID: id, ExpiredAt: expiredAt}

	return jwtData, errx.Nil()
}
