package directives

import (
	"context"
	"errors"

	customMiddleware "github/Martin-Martinez4/metube_backend/middleware"

	"github.com/99designs/gqlgen/graphql"
)

func Authorization(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {

	tokencookie := ctx.Value(customMiddleware.TokenCookieKey)

	// vlaidate token here

	if tokencookie == nil {
		// block calling the next resolver
		return nil, errors.New("Access denied")
	}

	// or let it pass through
	return next(ctx)
}
