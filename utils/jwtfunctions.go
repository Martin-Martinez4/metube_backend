package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type customClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func CreateJWT(id string, minutesValid time.Duration) (string, error) {

	claims := customClaims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(minutesValid * time.Minute).UnixMilli(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	JWT_SECRET := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString([]byte(JWT_SECRET))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Vlaidate and Parse JWT
func ValidateJWT(tokenString string) (*customClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {

		// verify algorithm used
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		JWT_SECRET := []byte(os.Getenv("JWT_SECRET"))

		return JWT_SECRET, nil

	})
	if err != nil {

		return nil, err
	}

	claims, ok := token.Claims.(*customClaims)

	if !ok {

		return nil, errors.New("token format error")
	}

	if !token.Valid {

		return nil, errors.New("token validation error")
	}

	return claims, nil

}
