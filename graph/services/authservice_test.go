package services

import (
	"context"
	db "github/Martin-Martinez4/metube_backend/config"
	"github/Martin-Martinez4/metube_backend/graph/model"
	helpers "github/Martin-Martinez4/metube_backend/testHelpers"
	"github/Martin-Martinez4/metube_backend/utils"
	"net/http/httptest"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestAuthServiceSQL_Login(t *testing.T) {

	TEST_DB_URL := db.ReadEnv("../../.env").TEST_DB_URL

	DB := db.GetDB("postgres", TEST_DB_URL)

	defer DB.Close()

	type args struct {
		login model.LoginInput
	}
	tests := []struct {
		name    string
		authsql *AuthServiceSQL
		args    args
		want    *model.Profile
		wantErr bool
	}{
		{
			name:    "login success",
			authsql: &AuthServiceSQL{DB: DB},
			args: args{
				login: model.LoginInput{
					Username: "coding_channel",
					Password: "password",
				},
			},
			want:    helpers.CodingChannelProfile,
			wantErr: false,
		},
		{
			name:    "password mismatch",
			authsql: &AuthServiceSQL{DB: DB},
			args: args{
				login: model.LoginInput{
					Username: "coding_channel",
					Password: "passwor",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "user not found",
			authsql: &AuthServiceSQL{DB: DB},
			args: args{
				login: model.LoginInput{
					Username: "notFound",
					Password: "password",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := helpers.MyContext{}

			w := httptest.NewRecorder()
			ctx2 := context.WithValue(ctx, utils.ResponseWriterKey, w)

			got, err := tt.authsql.Login(ctx2, tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthServiceSQL.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthServiceSQL.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Need to access a db from a docker file
func TestAuthServiceSQL_Register(t *testing.T) {

	TEST_DB_URL := db.ReadEnv("../../.env").TEST_DB_URL

	DB := db.GetDB("postgres", TEST_DB_URL)
	defer DB.Close()

	type args struct {
		profile model.RegisterInput
	}
	tests := []struct {
		name    string
		authsql *AuthServiceSQL
		args    args
		want    *model.Profile
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Successful Registration",
			authsql: &AuthServiceSQL{DB: DB},
			args: args{
				model.RegisterInput{
					Username:    "test_channel",
					Displayname: "test_channel",
					Password:    "Password",
					Password2:   "Password",
				},
			},
			want:    helpers.TestChannelProfile,
			wantErr: false,
		},
		{
			name:    "duplicated username",
			authsql: &AuthServiceSQL{DB: DB},
			args: args{
				model.RegisterInput{
					Username:    "coding_channel",
					Displayname: "test_channel",
					Password:    "Password",
					Password2:   "Password",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Mock context to be able to set the header in the function
			ctx := helpers.MyContext{}
			// Mock a ResponseWriter
			w := httptest.NewRecorder()
			ctx2 := context.WithValue(ctx, utils.ResponseWriterKey, w)

			got, err := tt.authsql.Register(ctx2, tt.args.profile)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthServiceSQL.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthServiceSQL.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
