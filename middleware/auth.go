package middleware

import (
	"context"
	"github/Martin-Martinez4/metube_backend/utils"
	"net/http"
)

// Middlewares
// Store TokenCookie in context
func WithTokenCookie() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			tokenCookie, err := req.Cookie("Auth")
			if err != nil {
				next.ServeHTTP(w, req)
				return
			}

			ctx := context.WithValue(req.Context(), utils.TokenCookieKey, tokenCookie.Value)

			next.ServeHTTP(w, req.WithContext(ctx))

		})
	}
}

// To be able to access ResponseWriter from the request conttext
func WithWriter() func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			ctx := req.Context()
			ctx = context.WithValue(ctx, utils.ResponseWriterKey, w)

			req = req.WithContext(ctx)
			next.ServeHTTP(w, req)

		})

	}
}
