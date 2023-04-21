package directives

import (
	"context"
	"errors"
	"strings"
	"time"

	utils "github/Martin-Martinez4/metube_backend/utils"

	"github.com/99designs/gqlgen/graphql"
)

type contextKey string

const UserKey = contextKey("user")

func Authorization(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {

	tokencookie := ctx.Value(utils.TokenCookieKey)
	if tokencookie == nil {
		// block calling the next resolver
		return nil, errors.New("no token present, access denied")
	}

	// validate token here
	tokenSlice := strings.Split(tokencookie.(string), "Bearer ")
	if len(tokenSlice) < 2 {

		return nil, errors.New("malformed token, access denied")

	}

	token := tokenSlice[1]

	claims, err := utils.ValidateJWT(token)
	if err != nil || claims == nil {

		return nil, errors.New("jwt validation error, access denied")
	}

	if claims.StandardClaims.ExpiresAt < time.Now().UnixMilli() {

		return nil, errors.New("token has expired")
	}

	idFromClaims := claims.Id
	if idFromClaims == "" {
		return nil, errors.New("claim id is empty. access denied")
	}

	newctx := context.WithValue(ctx, utils.UserKey, idFromClaims)

	return next(newctx)
}

func AuthorizationOptional(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {

	tokencookie := ctx.Value(utils.TokenCookieKey)
	if tokencookie == nil {
		// block calling the next resolver
		return next(ctx)
	}

	// validate token here
	tokenSlice := strings.Split(tokencookie.(string), "Bearer ")
	if len(tokenSlice) < 2 {

		return next(ctx)

	}

	token := tokenSlice[1]

	claims, err := utils.ValidateJWT(token)
	if err != nil || claims == nil {

		return next(ctx)
	}

	idFromClaims := claims.Id
	if idFromClaims == "" {
		return next(ctx)
	}

	newctx := context.WithValue(ctx, utils.UserKey, idFromClaims)

	return next(newctx)
}
