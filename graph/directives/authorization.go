package directives

import (
	"context"
	"errors"

	customMiddleware "github/Martin-Martinez4/metube_backend/middleware"

	"github.com/99designs/gqlgen/graphql"
)

type contextKey string

const UserKey = contextKey("user")

func Authorization(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {

	tokencookie := ctx.Value(customMiddleware.TokenCookieKey)
	if tokencookie == nil {
		// block calling the next resolver
		return nil, errors.New("access denied")
	}

	// validate token here
	claims, err := customMiddleware.ValidateJWT(tokencookie.(string))
	if err != nil || claims == nil {

		return nil, errors.New("access denied")
	}

	idFromClaims := claims["id"]
	if idFromClaims == "" {
		return nil, errors.New("access denied")
	}

	ctx = context.WithValue(ctx, UserKey, claims["id"])

	return next(ctx)
}
