package graphql

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"github/Martin-Martinez4/metube_backend/graph"
	"github/Martin-Martinez4/metube_backend/graph/model"
)

type Resolver struct{}

// // foo
func (r *commentResolver) Likes(ctx context.Context, obj *model.Comment) (int, error) {
	panic("not implemented")
}

// // foo
func (r *commentResolver) Dislikes(ctx context.Context, obj *model.Comment) (int, error) {
	panic("not implemented")
}

// // foo
func (r *commentResolver) Responses(ctx context.Context, obj *model.Comment) (int, error) {
	panic("not implemented")
}

// // foo
func (r *commentResolver) Profile(ctx context.Context, obj *model.Comment) (*model.Profile, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) UpsertVideo(ctx context.Context, input model.VideoInput) (*model.Video, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) Login(ctx context.Context, login model.LoginInput) (*model.Profile, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) Register(ctx context.Context, profileToRegister model.RegisterInput) (*model.Profile, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) Subscribe(ctx context.Context, subscribee string) (bool, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) Unsubscribe(ctx context.Context, subscribee string) (bool, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) VideoView(ctx context.Context, videoID string) (bool, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) LikeVideo(ctx context.Context, videoID string) (bool, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) DislikeVideo(ctx context.Context, videoID string) (bool, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) DeleteLikeDislikeVideo(ctx context.Context, videoID string) (bool, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) CreateComment(ctx context.Context, comment model.CommentInput) (*model.Comment, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) CreateResponse(ctx context.Context, comment model.CommentInput, parentCommentID string) (*model.Comment, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) LikeComment(ctx context.Context, commentID string) (bool, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) DislikeComment(ctx context.Context, commentID string) (bool, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) DeleteLikeDislikeComment(ctx context.Context, commentID string) (bool, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Videos(ctx context.Context, amount *int) ([]*model.Video, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Video(ctx context.Context, id string) (*model.Video, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) GetVideoLikeStatus(ctx context.Context, id string) (*model.LikeDislike, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) GetVideoComments(ctx context.Context, videoID string) ([]*model.Comment, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) GetCommentResponses(ctx context.Context, commentID string) ([]*model.Comment, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Profile(ctx context.Context, username string) (*model.Profile, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Profiles(ctx context.Context, amount int) ([]*model.Profile, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) GetLoggedInProfile(ctx context.Context) (*model.Profile, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) GetMentions(ctx context.Context) ([]*model.Comment, error) {
	panic("not implemented")
}

// // foo
func (r *videoResolver) Contentinformation(ctx context.Context, obj *model.Video) (*model.ContentInformation, error) {
	panic("not implemented")
}

// // foo
func (r *videoResolver) Thumbnail(ctx context.Context, obj *model.Video) (*model.Thumbnail, error) {
	panic("not implemented")
}

// // foo
func (r *videoResolver) Statistic(ctx context.Context, obj *model.Video) (*model.Statistic, error) {
	panic("not implemented")
}

// // foo
func (r *videoResolver) Status(ctx context.Context, obj *model.Video) (*model.Status, error) {
	panic("not implemented")
}

// // foo
func (r *videoResolver) Profile(ctx context.Context, obj *model.Video) (*model.Profile, error) {
	panic("not implemented")
}

// Comment returns graph.CommentResolver implementation.
func (r *Resolver) Comment() graph.CommentResolver { return &commentResolver{r} }

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// Video returns graph.VideoResolver implementation.
func (r *Resolver) Video() graph.VideoResolver { return &videoResolver{r} }

type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type videoResolver struct{ *Resolver }
