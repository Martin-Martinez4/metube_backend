package graph

import (
	"context"
	"database/sql"
	"github/Martin-Martinez4/metube_backend/graph/model"
	"log"
	"net/http"
	"time"
)

// go run github.com/vektah/dataloaden VideoLoader string *github/Martin-Martinez4/metube_backend/graph/model.Video

const videoloaderkey = "videoloader"

func DatatloaderMiddleware(db *sql.DB, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		videoloader := VideoLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*model.Video, []error) {

				var videos []*model.Video

				rows, err := db.Query("SELECT * FROM video WHERE id IN $1")
				defer rows.Close()
				if err != nil {
					return nil, []error{err}
				}

				for rows.Next() {
					video := model.Video{}

					err := rows.Scan(&video, &video)
					if err != nil {
						log.Fatal(err)
					}

					videos = append(videos, &video)

				}

				return videos, nil
			},
		}

		ctx := context.WithValue(req.Context(), videoloaderkey, &videoloader)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func getVideoLoader(ctx context.Context) *VideoLoader {

	return ctx.Value(videoloaderkey).(*VideoLoader)
}
