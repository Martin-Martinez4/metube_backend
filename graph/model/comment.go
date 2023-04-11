package model

type Comment struct {
	ID             string       `json:"id"`
	DatePosted     string       `json:"datePosted"`
	Body           string       `json:"body"`
	VideoID        *string      `json:"video_id"`
	ParentID       *string      `json:"parent_id"`
	Status         *LikeDislike `json:"status"`
	CommentProfile *Profile     `json:"CommentProfile"`
}
