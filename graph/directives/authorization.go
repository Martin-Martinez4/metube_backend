package directives

import (
	"context"
	"errors"
	"strings"

	utils "github/Martin-Martinez4/metube_backend/utils"

	"github.com/99designs/gqlgen/graphql"
)

type contextKey string

const UserKey = contextKey("user")

func Authorization(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {

	tokencookie := ctx.Value(utils.TokenCookieKey)
	if tokencookie == nil {
		// block calling the next resolver
		return nil, errors.New("access denied")
	}

	// validate token here
	token := strings.Split(tokencookie.(string), "Bearer ")[1]
	claims, err := utils.ValidateJWT(token)
	if err != nil || claims == nil {

		return nil, errors.New("access denied")
	}

	idFromClaims := claims["id"].(string)
	if idFromClaims == "" {
		return nil, errors.New("access denied")
	}

	// newctx := context.WithValue(ctx, "test", "test")
	newctx := context.WithValue(ctx, "user", idFromClaims)

	return next(newctx)
}
