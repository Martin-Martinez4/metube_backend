package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type contextKey string

const TokenCookieKey = contextKey("tokencookie")
const ResponseWriterKey = contextKey("responsewriter")

// Middlewares
// Store TokenCookie in context
func WithTokenCookie() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			tokenCookie, err := req.Cookie("token")
			if err != nil {
				next.ServeHTTP(w, req)
				return
			}

			ctx := context.WithValue(req.Context(), TokenCookieKey, tokenCookie.Value)

			next.ServeHTTP(w, req.WithContext(ctx))

		})
	}
}

// To be able to access ResponseWriter from the request conttext
func WithWriter() func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			ctx := req.Context()
			ctx = context.WithValue(ctx, ResponseWriterKey, w)

			req = req.WithContext(ctx)
			next.ServeHTTP(w, req)

		})

	}
}

func CreateJWT(id string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        id,
		"ExpiresAt": time.Now().Add(10 * time.Minute).UnixMilli(),
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
