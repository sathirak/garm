package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sathirak/garm/internal/config"
)

type JWT struct {
	ID        string
	ExpiredAt time.Time
}

func Generate(id string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(config.Get().App.JWTExpTime).Unix(),
	})

	tokenString, err := token.SignedString(GetKey())
	if err != nil {
		return "", err
	}

	bearerToken := "Bearer " + tokenString

	return bearerToken, nil
}

func Parse(bearerToken string) (jwtData *JWT, err error) {

	prefix := "Bearer "
	if !(len(bearerToken) > len(prefix) && bearerToken[:len(prefix)] == prefix) {
		return nil, fmt.Errorf("invalid token")
	}

	token, err := jwt.Parse(bearerToken[len(prefix):], func(token *jwt.Token) (interface{}, error) {
		return GetKey(), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	id := claims["id"].(string)
	exp := claims["exp"].(float64)
	expiredAt := time.Unix(int64(exp), 0)

	jwtData = &JWT{ID: id, ExpiredAt: expiredAt}

	return jwtData, nil
}
