package utils

import (
	"fmt"
	"github/Martin-Martinez4/metube_backend/config"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func TestCreateJWT(t *testing.T) {
	JWT_SECRET := config.ReadEnv("../.env").JWT_SECRET

	type args struct {
		id           string
		minutesValid time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "create valid jwt",
			args: args{
				id:           "testid",
				minutesValid: 10,
			},
			want:    "testid",
			wantErr: false,
		},
		{
			name: "id mismatch",
			args: args{
				id:           "",
				minutesValid: 10,
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateJWT(tt.args.id, tt.args.minutesValid)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			token, err := jwt.Parse(got, func(token *jwt.Token) (interface{}, error) {

				// verify algorithm used
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				JWT_SECRET := []byte(JWT_SECRET)

				return JWT_SECRET, nil

			})
			if err != nil {

				t.Errorf("token parsing error")
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok {

				t.Errorf("token format error")
			}

			if !token.Valid {

				t.Errorf("token validation error")
			}
			if claims["id"] != tt.want {
				t.Errorf("CreateJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}
