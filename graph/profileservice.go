package graph

import (
	"database/sql"
	"github/Martin-Martinez4/metube_backend/graph/model"
)

type ProfileService interface {
	GetProfileById(id string) (*model.Profile, error)
	GetMultipleProfiles(amount int) ([]*model.Profile, error)
}

type ProfileServiceSQL struct {
	DB *sql.DB
}

func (psql *ProfileServiceSQL) GetProfileById(id string) (*model.Profile, error) {

	row := psql.DB.QueryRow("SELECT id, username, displayname, isChannel, subscribers FROM profile WHERE id = $1", id)

	profile := model.Profile{}

	err := row.Scan(&profile.ID, &profile.Username, &profile.Displayname, &profile.IsChannel, &profile.Subscribers)

	if err != nil {

		return nil, err
	}

	return &profile, nil

}

func (psql *ProfileServiceSQL) GetMultipleProfiles(amount int) ([]*model.Profile, error) {

	rows, err := psql.DB.Query("SELECT id, username, displayname, isChannel, subscribers FROM profile ORDER BY RANDOM() LIMIT $1", amount)
	if err != nil {
		return nil, err
	}

	profiles := []*model.Profile{}

	for rows.Next() {

		profile := model.Profile{}

		err := rows.Scan(&profile.ID, &profile.Username, &profile.Displayname, &profile.IsChannel, &profile.Subscribers)

		if err != nil {
			return nil, err
		}

		profiles = append(profiles, &profile)
	}

	return profiles, nil

}
