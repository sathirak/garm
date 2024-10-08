package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

type JWTData struct {
	ID        string 
	ExpiredAt time.Time
}

func JWTGen(id string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JWTParse(tokenString string) (jwtData *JWTData, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
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

	jwtData = &JWTData{ID: id, ExpiredAt: expiredAt}
	fmt.Printf("Username: %s, Exp: %v\n", id, exp)

	return jwtData, nil
}
