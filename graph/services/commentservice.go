package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github/Martin-Martinez4/metube_backend/graph/model"
	"github/Martin-Martinez4/metube_backend/utils"
)

type CommentService interface {
	LikeComment(ctx context.Context, commentID string) (bool, error)
	DislikeComment(ctx context.Context, commentID string) (bool, error)
	DeleteLikeDislikeComment(ctx context.Context, commentID string) (bool, error)

	CreateComment(ctx context.Context, comment model.CommentInput) (bool, error)
	CreateResponse(ctx context.Context, comment model.CommentInput, parentCommentID string) (bool, error)

	GetVideoComments(ctx context.Context, videoID string) ([]*model.Comment, error)
	GetCommentResponses(ctx context.Context, commentID string) ([]*model.Comment, error)
	GetMentions(ctx context.Context) ([]*model.Comment, error)
}

type CommentServiceSQL struct {
	DB *sql.DB
}

func (csql *CommentServiceSQL) LikeComment(ctx context.Context, commentID string) (bool, error) {
	likerId := ctx.Value(utils.UserKey)
	if likerId == nil {
		return false, errors.New("token is nil")
	}

	_, err := csql.DB.Exec("INSERT INTO profile_comment_like_dislike(profile_id, comment_id, status) VALUES($1, $2, 'like') ON CONFLICT (profile_id, comment_id) DO UPDATE SET status = 'like'", likerId, commentID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (csql *CommentServiceSQL) DislikeComment(ctx context.Context, commentID string) (bool, error) {

	dislikerId := ctx.Value(utils.UserKey)
	if dislikerId == nil {
		return false, errors.New("token is nil")
	}

	_, err := csql.DB.Exec("INSERT INTO profile_comment_like_dislike(profile_id, comment_id, status) VALUES($1, $2, 'dislike') ON CONFLICT (profile_id, comment_id) DO UPDATE SET status = 'dislike'", dislikerId, commentID)
	if err != nil {
		return false, err
	}

	return true, nil

}

func (csql *CommentServiceSQL) DeleteLikeDislikeComment(ctx context.Context, commentID string) (bool, error) {

	profileId := ctx.Value(utils.UserKey)
	if profileId == nil {
		return false, errors.New("token is nil")
	}

	_, err := csql.DB.Exec("DELETE FROM profile_comment_like_dislike WHERE profile_id = $1 AND comment_id = $2", profileId, commentID)
	if err != nil {
		return false, err
	}

	return true, nil

}

func (csql *CommentServiceSQL) CreateComment(ctx context.Context, comment model.CommentInput) (bool, error) {

	commentorId := ctx.Value(utils.UserKey)
	if commentorId == nil {
		return false, errors.New("token is nil")
	}

	fail := func(err error) error {
		return fmt.Errorf("CreateComment: %v", err)
	}

	tx, err := csql.DB.BeginTx(ctx, nil)
	if err != nil {
		return false, fail(err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	row := tx.QueryRowContext(ctx, "INSERT INTO comment(id, date_posted, body, video_id, profile_id) VALUES (uuid_generate_v4(), now()::TIMESTAMPTZ, $1, $2, $3) RETURNING id", comment.Body, comment.VideoID, commentorId)
	if err != nil {
		return false, errors.New("inset comment failed")
	}

	var commentId string

	row.Scan(&commentId)

	valueStrings := []string{}
	valueArgs := []any{}
	i := 0

	// mentions in the body
	pattern := regexp.MustCompile(`\B\@([\w\-]+)`)
	mentions := pattern.FindAllString(comment.Body, 20)

	if len(mentions) > 0 {

		for _, mention := range mentions {
			valueStrings = append(valueStrings, fmt.Sprintf("$%d", i+2))
			valueArgs = append(valueArgs, fmt.Sprintf("%s", string(mention[1:])))
			i++

		}

		stmt := fmt.Sprintf("INSERT INTO profile_comment_mention(comment_id, profile_id) SELECT $1, id FROM profile WHERE username IN (%s)", strings.Join(valueStrings, ","))

		valueArgs = append([]any{commentId}, valueArgs...)
		for _, value := range valueArgs {

			fmt.Printf("%v \n", value)
		}

		_, err = tx.ExecContext(ctx, stmt, valueArgs...)
		if err != nil {
			return false, err
		}
	}

	if err = tx.Commit(); err != nil {
		return false, fail(err)
	}

	return true, nil

}

func (csql *CommentServiceSQL) CreateResponse(ctx context.Context, comment model.CommentInput, parentCommentID string) (bool, error) {

	// Handle setting parent comment id in the frontend

	commentorId := ctx.Value(utils.UserKey)
	if commentorId == nil {
		return false, errors.New("token is nil")
	}

	fail := func(err error) error {
		return fmt.Errorf("CreateComment: %v", err)
	}

	tx, err := csql.DB.BeginTx(ctx, nil)
	if err != nil {
		return false, fail(err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	row := tx.QueryRowContext(ctx, "INSERT INTO comment(id, date_posted, body, video_id, profile_id, parent_comment) VALUES (uuid_generate_v4(), now()::TIMESTAMPTZ, $1, $2, $3, $4) RETURNING id", comment.Body, comment.VideoID, commentorId, parentCommentID)
	if err != nil {
		return false, errors.New("inset comment failed " + err.Error())
	}

	var commentId string

	row.Scan(&commentId)

	valueStrings := []string{}
	valueArgs := []any{}
	i := 0

	pattern := regexp.MustCompile(`\B\@([\w\-]+)`)
	mentions := pattern.FindAllString(comment.Body, 20)

	for _, mention := range mentions {
		valueStrings = append(valueStrings, fmt.Sprintf("$%d", i+2))
		valueArgs = append(valueArgs, fmt.Sprintf("%s", string(mention[1:])))
		i++

	}

	stmt := fmt.Sprintf("INSERT INTO profile_comment_mention(comment_id, profile_id) SELECT $1, id FROM profile WHERE username IN (%s)", strings.Join(valueStrings, ","))

	valueArgs = append([]any{commentId}, valueArgs...)
	for _, value := range valueArgs {

		fmt.Printf("%v \n", value)
	}

	_, err = tx.ExecContext(ctx, stmt, valueArgs...)
	if err != nil {
		return false, err
	}

	if err = tx.Commit(); err != nil {
		return false, fail(err)
	}

	return true, nil

}

func (csql *CommentServiceSQL) GetVideoComments(ctx context.Context, videoID string) ([]*model.Comment, error) {

	profileId := ctx.Value(utils.UserKey)

	var rows *sql.Rows
	var err error

	if profileId != nil {

		rows, err = csql.DB.Query("SELECT id, date_posted, body, parent_comment, (SELECT status FROM profile_comment_like_dislike WHERE comment_id = comment.id AND profile_id = $1) FROM comment WHERE video_id = $2", profileId, videoID)
		if err != nil {
			return nil, err
		}

	} else {

		rows, err = csql.DB.Query("SELECT id, date_posted, body, parent_comment, NULL AS status FROM comment WHERE video_id = $1", videoID)
		if err != nil {
			return nil, err
		}

	}

	comments := []*model.Comment{}

	for rows.Next() {

		comment := model.Comment{}

		err := rows.Scan(&comment.ID, &comment.DatePosted, &comment.Body, &comment.ParentID, &comment.Status)
		if err != nil {
			return nil, err
		}

		comments = append(comments, &comment)

	}

	return comments, nil

}
func (csql *CommentServiceSQL) GetCommentResponses(ctx context.Context, commentID string) ([]*model.Comment, error) {
	profileId := ctx.Value(utils.UserKey)

	var rows *sql.Rows
	var err error

	if profileId != nil {

		rows, err = csql.DB.Query("SELECT id, date_posted, body, parent_comment, (SELECT status FROM profile_comment_like_dislike WHERE comment_id = comment.id AND profile_id = $1) FROM comment WHERE parent_comment = $2", profileId, commentID)
		if err != nil {
			return nil, err
		}

	} else {

		rows, err = csql.DB.Query("SELECT id, date_posted, body, parent_comment, NULL AS status FROM comment WHERE parent_comment = $1", commentID)
		if err != nil {
			return nil, err
		}

	}

	comments := []*model.Comment{}

	for rows.Next() {

		comment := model.Comment{}

		err := rows.Scan(&comment.ID, &comment.DatePosted, &comment.Body, &comment.ParentID, &comment.Status)
		if err != nil {
			return nil, err
		}

		comments = append(comments, &comment)

	}

	return comments, nil
}
func (csql *CommentServiceSQL) GetMentions(ctx context.Context) ([]*model.Comment, error) {

	profileId := ctx.Value(utils.UserKey)
	if profileId == nil {
		return nil, errors.New("access denied")
	}

	rows, err := csql.DB.Query("SELECT id, date_posted, body, parent_comment, video_id FROM comment JOIN profile_comment_mention ON profile_comment_mention.profile_id = $1", profileId)
	if err != nil {

		return nil, err
	}

	mentions := []*model.Comment{}

	for rows.Next() {

		mention := model.Comment{}

		rows.Scan(&mention.ID, &mention.DatePosted, &mention.Body, &mention.ParentID, &mention.VideoID)

		mentions = append(mentions, &mention)

	}

	return mentions, nil
}
