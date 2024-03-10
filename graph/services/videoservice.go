package services

import (
	"context"
	"database/sql"
	"github/Martin-Martinez4/metube_backend/graph/model"
	"github/Martin-Martinez4/metube_backend/utils"
)

type VideoService interface {
	GetVideoById(id string) (*model.Video, error)
<<<<<<< HEAD
=======
	SearchForVideoByTitle(searchTerm string) ([]*model.Video, error)
>>>>>>> master
	GetVideoLikeStatus(ctx context.Context, id string) (*model.LikeDislike, error)
	GetContentInformation(id string) (*model.ContentInformation, error)
	GetThumbnail(id string) (*model.Thumbnail, error)
	GetStatistic(id string) (*model.Statistic, error)
	GetStatus(id string) (*model.Status, error)
	GetProfile(ctx context.Context, id string) (*model.Profile, error)
	GetMultipleVideos(amount int) ([]*model.Video, error)
}

type VideoServiceSQL struct {
	DB *sql.DB
}

func (vsql *VideoServiceSQL) GetVideoById(id string) (*model.Video, error) {

	video := model.Video{}

	row := vsql.DB.QueryRow("SELECT id, url, categoryid, duration, profile_id FROM video WHERE id = $1", id)

	err := row.Scan(&video.ID, &video.URL, &video.Categoryid, &video.Duration, &video.ProfileID)
	if err != nil {
		return nil, err
	}

	return &video, nil
}

func (vsql *VideoServiceSQL) GetVideoLikeStatus(ctx context.Context, id string) (*model.LikeDislike, error) {

	profileId := ctx.Value(utils.UserKey)

	row := vsql.DB.QueryRow("SELECT status FROM profile_video_like_dislike WHERE profile_id = $1 AND video_id = $2", profileId, id)

	var status *model.LikeDislike

	err := row.Scan(&status)

	if err == sql.ErrNoRows {

		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	return status, nil
}

func (vsql *VideoServiceSQL) GetMultipleVideos(amount int) ([]*model.Video, error) {
	// Limit the amount
	rows, err := vsql.DB.Query("SELECT id, url, categoryid, duration, profile_id FROM video ORDER BY RANDOM() LIMIT $1", amount)
	if err != nil {
		return nil, err
	}

	videos := []*model.Video{}

	for rows.Next() {

		video := model.Video{}

		err := rows.Scan(&video.ID, &video.URL, &video.Categoryid, &video.Duration, &video.ProfileID)

		if err != nil {
			return nil, err
		}

		videos = append(videos, &video)
	}

	return videos, nil
}

func (vsql *VideoServiceSQL) GetContentInformation(id string) (*model.ContentInformation, error) {

	row := vsql.DB.QueryRow("SELECT title, description, channelid, published FROM contentinformation WHERE video_id = $1", id)

	contentinformation := model.ContentInformation{}

	err := row.Scan(&contentinformation.Title, &contentinformation.Description, &contentinformation.Channelid, &contentinformation.Published)
	if err != nil {
		return nil, err
	}

	return &contentinformation, nil
}

func (vsql *VideoServiceSQL) GetThumbnail(id string) (*model.Thumbnail, error) {
	row := vsql.DB.QueryRow("SELECT url FROM thumbnail WHERE video_id = $1", id)

	thumbnail := model.Thumbnail{}

	err := row.Scan(&thumbnail.URL)
	if err != nil {
		return nil, err
	}

	return &thumbnail, nil
}

func (vsql *VideoServiceSQL) GetStatistic(id string) (*model.Statistic, error) {
	row := vsql.DB.QueryRow("SELECT likes, dislikes, views, favorites, comments FROM statistic WHERE video_id = $1", id)

	statistic := model.Statistic{}

	err := row.Scan(&statistic.Likes, &statistic.Dislikes, &statistic.Views, &statistic.Favorites, &statistic.Comments)
	if err != nil {
		return nil, err
	}

	return &statistic, nil
}

func (vsql *VideoServiceSQL) GetStatus(id string) (*model.Status, error) {
	row := vsql.DB.QueryRow("SELECT uploadstatus, privacystatus FROM status WHERE video_id = $1", id)

	status := model.Status{}

	err := row.Scan(&status.Uploadstatus, &status.Privacystatus)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

func (vsql *VideoServiceSQL) GetProfile(ctx context.Context, id string) (*model.Profile, error) {

	profileId := ctx.Value(utils.UserKey)

	var row *sql.Row

	if profileId == nil {
		row = vsql.DB.QueryRow("SELECT username, displayname, isChannel, subscribers, false AS user_subscribed FROM profile WHERE id = $1", id)
	} else {

		row = vsql.DB.QueryRow("SELECT username, displayname, isChannel, subscribers, EXISTS(SELECT 1 FROM subscriber_subscribee WHERE subscriber_id = $1 AND subscribee_id = profile.id) AS user_subscribed FROM profile WHERE id = $2", profileId, id)
	}

	profile := model.Profile{}

	err := row.Scan(&profile.Username, &profile.Displayname, &profile.IsChannel, &profile.Subscribers, &profile.UserIsSubscribedTo)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
<<<<<<< HEAD
=======

func (vsql *VideoServiceSQL) SearchForVideoByTitle(searchTerm string) ([]*model.Video, error) {

	var similarityThershold float32 = 0.09

	// 	SELECT
	// 	title
	// FROM contentinformation
	// WHERE similarity(title, 'JavaScript') > .09;

	rows, err := vsql.DB.Query("SELECT id, url, categoryid, duration, profile_id FROM contentinformation JOIN video ON contentinformation.video_id = video.id WHERE similarity(contentinformation.title, $1) > $2", searchTerm, similarityThershold)
	if err != nil {
		return nil, err
	}

	videos := []*model.Video{}

	for rows.Next() {

		video := model.Video{}

		err := rows.Scan(&video.ID, &video.URL, &video.Categoryid, &video.Duration, &video.ProfileID)

		if err != nil {
			return nil, err
		}

		videos = append(videos, &video)
	}

	return videos, nil
}
>>>>>>> master
