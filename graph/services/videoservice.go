package services

import (
	"context"
	"database/sql"
	"fmt"
	"github/Martin-Martinez4/metube_backend/graph/model"
	"github/Martin-Martinez4/metube_backend/utils"

	"github.com/lib/pq"

	"github.com/graph-gophers/dataloader/v7"
)

type VideoService interface {
	GetVideoById(id string) (*model.Video, error)
	SearchForVideoByTitle(searchTerm string) ([]*model.Video, error)
	GetVideoLikeStatus(ctx context.Context, id string) (*model.LikeDislike, error)
	GetContentInformation(id string) (*model.ContentInformation, error)
	GetThumbnail(id string) (*model.Thumbnail, error)
	GetStatistic(id string) (*model.Statistic, error)
	GetStatus(id string) (*model.Status, error)
	GetProfile(ctx context.Context, id string) (*model.Profile, error)
	GetVideosByProfileUsername(profileUsername string) ([]*model.Video, error)
	GetMultipleVideos(amount int) ([]*model.Video, error)
	GetMultipleVideosSetOrder(ctx context.Context, seed *float64, limit *int, offset *int) ([]*model.Video, error)
}

type VideoServiceSQL struct {
	DB *sql.DB
}

func mapResults[K comparable, V any](keys []K, data map[K]V) []*dataloader.Result[V] {
	results := make([]*dataloader.Result[V], len(keys))

	for i, key := range keys {
		if val, ok := data[key]; ok {
			results[i] = &dataloader.Result[V]{Data: val}
		} else {
			results[i] = &dataloader.Result[V]{Data: *new(V)}
		}
	}

	return results
}

func (vsql *VideoServiceSQL) GetVideoById(id string) (*model.Video, error) {

	video := model.Video{}

	row := vsql.DB.QueryRow("SELECT id, url, categoryid, duration, profile_id FROM video WHERE id = $1 AND visible = TRUE", id)

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
	rows, err := vsql.DB.Query("SELECT id, url, categoryid, duration, profile_id FROM video WHERE visible = TRUE ORDER BY RANDOM() LIMIT $1", amount)
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

func (vsql *VideoServiceSQL) GetMultipleVideosSetOrder(ctx context.Context, seed *float64, limit *int, offset *int) ([]*model.Video, error) {
	// Limit the amount
	tx, err := vsql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// BEGIN; SELECT setseed(0.5); SELECT id, url, categoryid, duration, profile_id FROM video ORDER BY RANDOM() LIMIT 5; COMMIT;
	_, err = vsql.DB.Exec("SELECT setseed($1)", seed)
	if err != nil {
		return nil, err
	}
	_, err = vsql.DB.Exec(`SET pg_trgm.similarity_threshold = 0.05;`)
	if err != nil {
		return nil, err
	}
	rows, _ := vsql.DB.Query("SELECT id, url, categoryid, duration, profile_id FROM video WHERE visible = TRUE ORDER BY RANDOM() LIMIT $1 OFFSET $2", limit, offset)
	defer rows.Close()

	videos := []*model.Video{}

	fmt.Println(rows)
	for rows.Next() {

		video := model.Video{}

		err := rows.Scan(&video.ID, &video.URL, &video.Categoryid, &video.Duration, &video.ProfileID)

		if err != nil {
			return nil, err
		}

		videos = append(videos, &video)
	}

	if err = tx.Commit(); err != nil {
		return nil, err
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

func (vsql *VideoServiceSQL) BatchContentInformation(ctx context.Context, keys []string) []*dataloader.Result[*model.ContentInformation] {
	query := `
		SELECT video_id, title, description, channelid, published 
		FROM contentinformation 
		WHERE video_id = ANY($1)
	`

	rows, err := vsql.DB.QueryContext(ctx, query, pq.Array(keys))
	if err != nil {
		results := make([]*dataloader.Result[*model.ContentInformation], len(keys))
		for i := range results {
			results[i] = &dataloader.Result[*model.ContentInformation]{Error: err}
		}

		return results
	}
	defer rows.Close()

	contentMap := make(map[string]*model.ContentInformation)

	for rows.Next() {
		var videoID string
		var title string
		var description string
		var channelID string
		var published string

		if err := rows.Scan(&videoID, &title, &description, &channelID, &published); err != nil {
			continue
		}

		contentMap[videoID] = &model.ContentInformation{
			Title:     title,
			Channelid: channelID,
			Published: published,
		}

	}

	return mapResults(keys, contentMap)

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

func (vsql *VideoServiceSQL) BatchGetThumbnail(ctx context.Context, keys []string) []*dataloader.Result[*model.Thumbnail] {
	query := "SELECT video_id, url FROM thumbnail WHERE video_id = ANY($1)"
	rows, err := vsql.DB.QueryContext(ctx, query, pq.Array(keys))
	if err != nil {
		results := make([]*dataloader.Result[*model.Thumbnail], len(keys))

		for i := range results {
			results[i] = &dataloader.Result[*model.Thumbnail]{Error: err}
		}

		return results
	}
	defer rows.Close()

	thumbnailMap := make(map[string]*model.Thumbnail)

	for rows.Next() {
		var videoID string
		var url string

		if err := rows.Scan(&videoID, &url); err != nil {
			continue
		}

		thumbnailMap[videoID] = &model.Thumbnail{
			URL: url,
		}
	}

	return mapResults(keys, thumbnailMap)

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

func (vsql *VideoServiceSQL) BatchGetStatistic(ctx context.Context, keys []string) []*dataloader.Result[*model.Statistic] {
	query := "SELECT video_id, likes, dislikes, views, favorites, comments FROM statistic WHERE video_id = ANY($1)"
	rows, err := vsql.DB.QueryContext(ctx, query, pq.Array(keys))
	if err != nil {
		results := make([]*dataloader.Result[*model.Statistic], len(keys))
		for i := range results {
			results[i] = &dataloader.Result[*model.Statistic]{Error: err}
		}

		return results
	}
	defer rows.Close()

	statsMap := make(map[string]*model.Statistic)

	for rows.Next() {
		var videoID string
		var likes, dislikes, views, favorites, comments int

		if err := rows.Scan(&videoID, &likes, &dislikes, &views, &favorites, &comments); err != nil {
			continue
		}
		statsMap[videoID] = &model.Statistic{
			Likes:     likes,
			Dislikes:  dislikes,
			Views:     views,
			Favorites: &favorites,
			Comments:  comments,
		}

	}

	return mapResults(keys, statsMap)

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

func (vsql *VideoServiceSQL) SearchForVideoByTitle(searchTerm string) ([]*model.Video, error) {

	// var similarityThershold float32 = 0.09

	// 	SELECT
	// 	title
	// FROM contentinformation
	// WHERE similarity(title, 'JavaScript') > .09;

	rows, err := vsql.DB.Query(`
		SELECT v.id, v.url, v.categoryid, v.duration, v.profile_id
		FROM contentinformation ci
		JOIN video v ON ci.video_id = v.id
		WHERE ci.title % $1
		  AND v.visible = TRUE
		ORDER BY similarity(ci.title, $1) DESC
	`, searchTerm)
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

func (vsql *VideoServiceSQL) GetVideosByProfileUsername(profileUsername string) ([]*model.Video, error) {

	rows, err := vsql.DB.Query("SELECT video.id, url, categoryid, duration, profile_id FROM video JOIN profile ON profile.id = video.profile_id WHERE profile.username =$1 AND video.visible=true", profileUsername)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
