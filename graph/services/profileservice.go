package services

import (
	"context"
	"database/sql"
	"errors"
	"github/Martin-Martinez4/metube_backend/graph/model"
	"github/Martin-Martinez4/metube_backend/utils"
)

type ProfileService interface {
	GetProfileIdFromUsername(username string) (string, error)
	GetProfileByUsername(ctx context.Context, username string) (*model.Profile, error)
	GetMultipleProfiles(amount int) ([]*model.Profile, error)

	Subscribe(ctx context.Context, subscribee string) (bool, error)
	Unsubscribe(ctx context.Context, subscribee string) (bool, error)
	VideoView(ctx context.Context, videoID string) (bool, error)
	LikeVideo(ctx context.Context, videoID string) (bool, error)
	DislikeVideo(ctx context.Context, videoID string) (bool, error)
	DeleteLikeDislikeVideo(ctx context.Context, videoID string) (bool, error)
}

type ProfileServiceSQL struct {
	DB *sql.DB
}

func (psql *ProfileServiceSQL) GetProfileIdFromUsername(username string) (string, error) {

	row := psql.DB.QueryRow("SELECT id FROM profile WHERE username = $1", username)

	var id string

	err := row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil

}
func (psql *ProfileServiceSQL) GetProfileByUsername(ctx context.Context, username string) (*model.Profile, error) {

	profileId := ctx.Value(utils.UserKey)

	var row *sql.Row

	if profileId == nil {

		row = psql.DB.QueryRow("SELECT username, displayname, isChannel, subscribers, false FROM profile WHERE username = $1", username)
	} else {
		row = psql.DB.QueryRow("SELECT username, displayname, isChannel, subscribers, EXISTS(SELECT 1 FROM subscriber_subscribee WHERE subscriber_id = $1 AND subscribee_id = profile.id) AS user_subscribed FROM profile WHERE username = $2", profileId, username)

	}

	profile := model.Profile{}

	err := row.Scan(&profile.Username, &profile.Displayname, &profile.IsChannel, &profile.Subscribers, &profile.UserIsSubscribedTo)

	if err != nil {

		return nil, err
	}

	return &profile, nil

}

func (psql *ProfileServiceSQL) GetMultipleProfiles(amount int) ([]*model.Profile, error) {

	rows, err := psql.DB.Query("SELECT username, displayname, isChannel, subscribers FROM profile ORDER BY RANDOM() LIMIT $1", amount)
	if err != nil {
		return nil, err
	}

	profiles := []*model.Profile{}

	for rows.Next() {

		profile := model.Profile{}

		err := rows.Scan(&profile.Username, &profile.Displayname, &profile.IsChannel, &profile.Subscribers)

		if err != nil {
			return nil, err
		}

		profiles = append(profiles, &profile)
	}

	return profiles, nil

}

func (psql *ProfileServiceSQL) Subscribe(ctx context.Context, subscribee string) (bool, error) {

	subscriberId := ctx.Value(utils.UserKey)
	if subscriberId == nil {
		return false, errors.New("token is nil")
	}

	_, err := psql.DB.Exec("INSERT INTO subscriber_subscribee(subscriber_id, subscribee_id) SELECT $1, id FROM profile WHERE username = $2 ON CONFLICT DO NOTHING", subscriberId, subscribee)
	if err != nil {
		return false, err
	}

	return true, nil
}
func (psql *ProfileServiceSQL) Unsubscribe(ctx context.Context, subscribee string) (bool, error) {

	subscriberId := ctx.Value(utils.UserKey)
	if subscriberId == nil {
		return false, errors.New("token is nil")
	}

	_, err := psql.DB.Exec("DELETE FROM subscriber_subscribee WHERE subscriber_id = $1 AND subscribee_id = (SELECT id FROM profile WHERE username = $2)", subscriberId, subscribee)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (psql *ProfileServiceSQL) VideoView(ctx context.Context, videoID string) (bool, error) {

	viewerId := ctx.Value(utils.UserKey)
	if viewerId == nil {
		return false, errors.New("token is nil")
	}

	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO profile_view(profile_id, video_id) VALUES($1, $2) ON CONFLICT DO NOTHING", viewerId, videoID)
	if err != nil {
		return false, err
	}

	_, err = tx.Exec("UPDATE video SET views = views + 1 WHERE id = $1", videoID)
	if err != nil {
		return false, err
	}

	if err = tx.Commit(); err != nil {
		return false, err
	} else {

		return true, nil
	}

}

func (psql *ProfileServiceSQL) LikeVideo(ctx context.Context, videoID string) (bool, error) {

	likerId := ctx.Value(utils.UserKey)
	if likerId == nil {
		return false, errors.New("token is nil")
	}

	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	row := tx.QueryRow("SELECT status FROM profile_video_like_dislike WHERE profile_id = $1 AND video_id = $2", likerId, videoID)

	var currentStatus string

	err = row.Scan(&currentStatus)
	if err != nil && err != sql.ErrNoRows {

		return false, err
	}

	if currentStatus != "like" {

		_, err = tx.ExecContext(ctx, "INSERT INTO profile_video_like_dislike(profile_id, video_id, status) VALUES($1, $2, 'like') ON CONFLICT (profile_id, video_id) DO UPDATE SET status = 'like'", likerId, videoID)
		if err != nil {
			return false, err
		}

		if currentStatus == "dislike" {

			_, err = tx.ExecContext(ctx, "UPDATE statistic SET dislikes = dislikes - 1, likes = likes + 1 WHERE video_id = $1", videoID)
			if err != nil {
				return false, err
			}
		} else {
			_, err = tx.ExecContext(ctx, "UPDATE statistic SET likes = likes + 1 WHERE video_id = $1", videoID)
			if err != nil {
				return false, err
			}

		}

	}

	if err = tx.Commit(); err != nil {
		return false, err
	} else {

		return true, nil
	}

}

func (psql *ProfileServiceSQL) DislikeVideo(ctx context.Context, videoID string) (bool, error) {
	dislikerId := ctx.Value(utils.UserKey)
	if dislikerId == nil {
		return false, errors.New("token is nil")
	}

	tx, err := psql.DB.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	row := tx.QueryRow("SELECT status FROM profile_video_like_dislike WHERE profile_id = $1 AND video_id = $2", dislikerId, videoID)

	var currentStatus string

	err = row.Scan(&currentStatus)
	if err != nil && err != sql.ErrNoRows {

		return false, err
	}

	if currentStatus != "dislike" {

		_, err = tx.ExecContext(ctx, "INSERT INTO profile_video_like_dislike(profile_id, video_id, status) VALUES($1, $2, 'dislike') ON CONFLICT (profile_id, video_id) DO UPDATE SET status = 'dislike'", dislikerId, videoID)
		if err != nil {
			return false, err
		}

		if currentStatus == "like" {

			_, err = tx.ExecContext(ctx, "UPDATE statistic SET likes = likes - 1, dislikes = dislikes + 1 WHERE video_id = $1", videoID)
			if err != nil {
				return false, err
			}
		} else {

			_, err = tx.ExecContext(ctx, "UPDATE statistic SET dislikes = dislikes + 1 WHERE video_id = $1", videoID)
			if err != nil {
				return false, err
			}
		}

	}

	if err = tx.Commit(); err != nil {
		return false, err
	} else {

		return true, nil
	}
}

func (psql *ProfileServiceSQL) DeleteLikeDislikeVideo(ctx context.Context, videoID string) (bool, error) {
	profileId := ctx.Value(utils.UserKey)
	if profileId == nil {
		return false, errors.New("token is nil")
	}

	_, err := psql.DB.Exec("DELETE FROM profile_video_like_dislike WHERE profile_id = $1 AND video_id = $2", profileId, videoID)
	if err != nil {
		return false, err
	}

	return true, nil
}
