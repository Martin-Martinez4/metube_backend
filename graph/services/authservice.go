package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github/Martin-Martinez4/metube_backend/graph/model"
	"github/Martin-Martinez4/metube_backend/utils"

	"net/http"
	"time"

	bcrypt "golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, login model.LoginInput) (*model.Profile, error)
	Register(ctx context.Context, profile model.RegisterInput) (*model.Profile, error)
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
	err = bcrypt.CompareHashAndPassword([]byte(passwordFromDB), []byte(login.Password))
	if err != nil {

		return nil, errors.New("passwords do not mach")
	}

	jwtToken, err := utils.CreateJWT(idFromDB, 10)
	if err != nil {
		return nil, err

	}

	cookie := http.Cookie{
		Name:     "Auth",
		Value:    "Bearer " + jwtToken,
		Expires:  time.Now().Add(10 * time.Minute),
		HttpOnly: true,
	}

	// the writer is saved to the context by the use of the WithWriter() middleware located at middleware/auth.go/
	// It is used here to send the JWT as a cookie
	http.SetCookie(ctx.Value(utils.ResponseWriterKey).(http.ResponseWriter), &cookie)

	return profile, nil

}

func (authsql *AuthServiceSQL) Register(ctx context.Context, profile model.RegisterInput) (*model.Profile, error) {

	// Create a helper function for preparing failure results.
	fail := func(err error) error {
		return fmt.Errorf("CreateOrder: %v", err)
	}

	if profile.Password != profile.Password2 {

		return nil, errors.New("passwords must match")
	}

	// Get a Tx for making transaction requests.
	tx, err := authsql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, fail(err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// Check if user already exists
	var userExists bool
	if err = tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM profile WHERE username = $1)", profile.Username).Scan(&userExists); err != nil || userExists {

		return nil, fail(errors.New("choose another username"))
	}

	// insert into profile and get profile
	row := tx.QueryRowContext(ctx, "INSERT INTO profile(id, username, displayname, isChannel, subscribers) VALUES (uuid_generate_v4(), $1, $2, false, 0) RETURNING id, username, displayname, isChannel, subscribers", profile.Username, profile.Displayname)

	var id string
	profileToReturn := &model.Profile{}
	err = row.Scan(&id, &profileToReturn.Username, &profileToReturn.Displayname, &profileToReturn.IsChannel, &profileToReturn.Subscribers)
	if err != nil || profileToReturn.Username == "" || id == "" {
		return nil, fail(errors.New("inserting profile"))
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(profile.Password), 10)
	if err != nil {

		return nil, fail(err)
	}

	// insert into login
	_, err = tx.ExecContext(ctx, "INSERT INTO login(profile_id, password) VALUES ($1, $2)", id, string(passwordHash))
	if err != nil {
		return nil, fail(err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fail(err)
	}

	jwtToken, err := utils.CreateJWT(id, 10)
	if err != nil {
		return nil, err

	}

	cookie := http.Cookie{
		Name:     "Auth",
		Value:    "Bearer " + jwtToken,
		Expires:  time.Now().Add(10 * time.Minute),
		HttpOnly: true,
	}

	// the writer is saved to the context by the use of the WithWriter() middleware located at middleware/auth.go/
	// It is used here to send the JWT as a cookie
	http.SetCookie(ctx.Value(utils.ResponseWriterKey).(http.ResponseWriter), &cookie)

	return profileToReturn, nil

}
