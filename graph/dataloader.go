package graph

import (
	"context"
	"database/sql"
	"fmt"
	"github/Martin-Martinez4/metube_backend/graph/model"
	"github/Martin-Martinez4/metube_backend/graph/services"
	"net/http"
	"time"

	"github.com/graph-gophers/dataloader/v7"
)

// go run github.com/vektah/dataloaden VideoLoader string *github/Martin-Martinez4/metube_backend/graph/model.Video

const profileLoaderKey = "profileloader"
const loadersKey = "loadersKey"

type Loaders struct {
	ContentLoader    *dataloader.Loader[string, *model.ContentInformation]
	ThumbnailLoader  *dataloader.Loader[string, *model.Thumbnail]
	StatisticsLoader *dataloader.Loader[string, *model.Statistic]
}

func NewLoaders(db *sql.DB) *Loaders {
	reader := &services.VideoServiceSQL{DB: db}
	return &Loaders{
		ContentLoader: dataloader.NewBatchedLoader(
			reader.BatchContentInformation,
		),
		ThumbnailLoader: dataloader.NewBatchedLoader(
			reader.BatchGetThumbnail,
		),
		StatisticsLoader: dataloader.NewBatchedLoader(
			reader.BatchGetStatistic,
		),
	}
}

func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

func DatatloaderMiddleware(db *sql.DB, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, res *http.Request) {
		profileLoader := ProfileLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(keys []string) ([]*model.Profile, []error) {
				// This is where the batching happens

				type profileWithCommentId struct {
					model.Profile
					comment_id string
				}
				var profiles = make(map[string]*model.Profile, len(keys))

				var test string

				for index, _ := range keys {

					if test != "" {

						test = fmt.Sprintf("%s, $%v", test, index+1)
					} else {

						test = fmt.Sprintf("$%v", index+1)
					}

				}

				var tempKeys []any

				for _, value := range keys {

					tempKeys = append(tempKeys, value)

				}

				sql := fmt.Sprintf("SELECT comment.id, username, displayname, isChannel, subscribers, false AS user_subscribed FROM profile INNER JOIN (SELECT profile_id, id FROM comment WHERE comment.id IN (%s)) AS comment ON profile.id = comment.profile_id", test)

				rows, err := db.Query(sql, tempKeys...)

				if err != nil {
					return nil, []error{err}
				}

				for rows.Next() {
					var profile profileWithCommentId
					_ = rows.Scan(&profile.comment_id, &profile.Username, &profile.Displayname, &profile.IsChannel, &profile.Subscribers, &profile.UserIsSubscribedTo)

					profiles[profile.comment_id] = &model.Profile{
						Username:           profile.Username,
						Displayname:        profile.Displayname,
						IsChannel:          profile.IsChannel,
						Subscribers:        profile.Subscribers,
						UserIsSubscribedTo: profile.UserIsSubscribedTo,
					}
				}

				result := make([]*model.Profile, len(keys))

				for index, comment_id := range keys {
					result[index] = profiles[comment_id]
				}

				return result, []error{err}
			},
		}

		ctx := context.WithValue(res.Context(), profileLoaderKey, &profileLoader)

		loaders := NewLoaders(db)
		ctx = context.WithValue(ctx, loadersKey, loaders)

		next.ServeHTTP(w, res.WithContext(ctx))
	})

}

func GetProfileLoader(ctx context.Context) *ProfileLoader {
	return ctx.Value(profileLoaderKey).(*ProfileLoader)
}
