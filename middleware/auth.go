package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	services "github/Martin-Martinez4/metube_backend/graph/services"

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

func CreateJWT(username string, ps services.ProfileService) (string, error) {

	id, err := ps.GetProfileIdFromUsername(username)
	if err != nil {

		return "", err
	}

	// For now set the profile id later change to set the session id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        time.Now().Add(10 * time.Minute),
		"id":         id,
		"authorized": true,
	})

	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// verify algorithm used
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		JWT_SECRET := os.Getenv("JWT_SECRET")

		hmacSampleSecret := []byte(JWT_SECRET)
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil

	})
	if err != nil {

		return false, err
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {

		return false, errors.New("token validation error")
	}

	return true, nil

}
