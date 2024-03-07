package model

type Comment struct {
	ID             string       `json:"id"`
	DatePosted     string       `json:"datePosted"`
	Body           string       `json:"body"`
	VideoID        *string      `json:"video_id"`
	ParentID       *string      `json:"parent_id"`
	Likes          int          `json:"likes"`
	Dislikes       int          `json:"dislikes"`
	Responses      int          `json:"responses"`
	Status         *LikeDislike `json:"status"`
	CommentProfile *Profile     `json:"CommentProfile"`
}
