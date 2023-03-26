package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(id string, minutesValid time.Duration) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        id,
		"ExpiresAt": time.Now().Add(minutesValid * time.Minute).UnixMilli(),
	})

	JWT_SECRET := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString([]byte(JWT_SECRET))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Vlaidate and Parse JWT
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

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

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {

		return nil, errors.New("token format error")
	}

	if !token.Valid {

		return nil, errors.New("token validation error")
	}

	return claims, nil

}
