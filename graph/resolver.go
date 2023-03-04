package graph

import (
	"database/sql"
	"github/Martin-Martinez4/metube_backend/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type VideoRepo struct {
	DB *sql.DB
}

// psql or other server would go here
type Resolver struct {
	VideoStore map[string]model.Video
	VideoRepo  VideoRepo
}
