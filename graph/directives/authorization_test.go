package directives

import (
	"context"
	testHelpers "github/Martin-Martinez4/metube_backend/testHelpers"
	"github/Martin-Martinez4/metube_backend/utils"
	"reflect"
	"testing"

	"github.com/99designs/gqlgen/graphql"
)

// called by next in the auth directive, used to compare context value
func Resolver(ctx context.Context) (res interface{}, err error) {

	user := ctx.Value(utils.UserKey)

	return user, err

}

func TestAuthorization(t *testing.T) {

	validJWT, err := utils.CreateJWT(testHelpers.Anime_channel_id)
	if err != nil {

		t.Errorf("error creating jwt token")

	}
	emptyIdJWT, err2 := utils.CreateJWT("")
	if err2 != nil {

		t.Errorf("error creating jwt token")

	}

	myContext := testHelpers.MyContext{}
	validTokenContext := context.WithValue(myContext, utils.TokenCookieKey, "Bearer "+validJWT)
	bearMissingContext := context.WithValue(myContext, utils.TokenCookieKey, validJWT)
	invalidJWTContext := context.WithValue(myContext, utils.TokenCookieKey, "Bearer testInvalid")
	emptyIdContext := context.WithValue(myContext, utils.TokenCookieKey, "Bearer "+emptyIdJWT)

	type args struct {
		ctx  context.Context
		obj  interface{}
		next graphql.Resolver
	}
	tests := []struct {
		name           string
		args           args
		wantRes        interface{}
		wantErr        bool
		wantErrMessage string
	}{
		{
			name: "happy path, valid jwt, well formed cookie",
			args: args{
				ctx:  validTokenContext,
				obj:  0,
				next: Resolver,
			},
			wantRes:        testHelpers.Anime_channel_id,
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "bearer missing",
			args: args{
				ctx:  bearMissingContext,
				obj:  0,
				next: Resolver,
			},
			wantRes:        nil,
			wantErr:        true,
			wantErrMessage: "malformed token, access denied",
		},
		{
			name: "invalid jwt",
			args: args{
				ctx:  invalidJWTContext,
				obj:  0,
				next: Resolver,
			},
			wantRes:        nil,
			wantErr:        true,
			wantErrMessage: "jwt validation error, access denied",
		},
		{
			name: "jwt with empty claim[id]",
			args: args{
				ctx:  emptyIdContext,
				obj:  0,
				next: Resolver,
			},
			wantRes:        nil,
			wantErr:        true,
			wantErrMessage: "claim id is empty. access denied",
		},
		{
			name: "jwt missing",
			args: args{
				ctx:  myContext,
				obj:  0,
				next: Resolver,
			},
			wantRes:        nil,
			wantErr:        true,
			wantErrMessage: "no token present, access denied",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := Authorization(tt.args.ctx, tt.args.obj, tt.args.next)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authorization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) && tt.wantErr {
				if err.Error() != tt.wantErrMessage {
					t.Errorf("Authorization() error = %v, wantErr %v", err.Error(), tt.wantErrMessage)
					return
				}
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Authorization() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
