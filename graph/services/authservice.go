package services

import (
	"context"
	"database/sql"
	"errors"
	"github/Martin-Martinez4/metube_backend/graph/model"
	middlewares "github/Martin-Martinez4/metube_backend/middleware"
	"net/http"
	"time"
)

type AuthService interface {
	Login(ctx context.Context, login model.LoginInput) (*model.Profile, error)
}

type AuthServiceSQL struct {
	DB *sql.DB
}

func (authsql *AuthServiceSQL) Login(ctx context.Context, login model.LoginInput) (*model.Profile, error) {

	// Validate LoginInput
	row := authsql.DB.QueryRow("SELECT id, password, username, displayname, isChannel, subscribers FROM profile JOIN login ON profile.id = login.profile_id WHERE profile.username = $1", login.Username)

	var passwordFromDB string
	var idFromDB string

	profile := &model.Profile{}
	err := row.Scan(&idFromDB, &passwordFromDB, &profile.Username, &profile.Displayname, &profile.IsChannel, &profile.Subscribers)
	if err != nil {
		return nil, err
	}
	// Change to bcrypt later
	if passwordFromDB != login.Password {

		return nil, errors.New("passwords do not mach")
	}

	jwtToken, err := middlewares.CreateJWT(idFromDB)
	if err != nil {
		return nil, err

	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    "Bearer " + jwtToken,
		Expires:  time.Now().Add(10 * time.Minute),
		HttpOnly: true,
	}

	// the writer is saved to the context by the use of the WithWriter() middleware located at middleware/auth.go/
	// It is used here to send the JWT as a cookie
	http.SetCookie(ctx.Value(middlewares.ResponseWriterKey).(http.ResponseWriter), &cookie)

	return profile, nil

}
