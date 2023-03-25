package directives

import (
	"context"
	"reflect"
	"testing"

	"github.com/99designs/gqlgen/graphql"
)

func TestAuthorization(t *testing.T) {
	type args struct {
		ctx  context.Context
		obj  interface{}
		next graphql.Resolver
	}
	tests := []struct {
		name    string
		args    args
		wantRes interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := Authorization(tt.args.ctx, tt.args.obj, tt.args.next)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authorization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Authorization() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
