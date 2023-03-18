package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"
	"github/Martin-Martinez4/metube_backend/graph/model"
)

// UpsertVideo is the resolver for the upsertVideo field.
func (r *mutationResolver) UpsertVideo(ctx context.Context, input model.VideoInput) (*model.Video, error) {
	panic(fmt.Errorf("not implemented: UpsertVideo - upsertVideo"))
}

// Videos is the resolver for the videos field.
func (r *queryResolver) Videos(ctx context.Context, amount *int) ([]*model.Video, error) {
	return r.VideoService.GetMultipleVideos(*amount)
}

// Video is the resolver for the video field.
func (r *queryResolver) Video(ctx context.Context, id string) (*model.Video, error) {
	return r.VideoService.GetVideoById(id)
}

// Contentinformation is the resolver for the contentinformation field.
func (r *videoResolver) Contentinformation(ctx context.Context, obj *model.Video) (*model.ContentInformation, error) {
	return r.VideoService.GetContentInformation(obj.ID)
}

// Thumbnail is the resolver for the thumbnail field.
func (r *videoResolver) Thumbnail(ctx context.Context, obj *model.Video) (*model.Thumbnail, error) {
	return r.VideoService.GetThumbnail(obj.ID)
}

// Statistic is the resolver for the statistic field.
func (r *videoResolver) Statistic(ctx context.Context, obj *model.Video) (*model.Statistic, error) {
	return r.VideoService.GetStatistic(obj.ID)
}

// Status is the resolver for the status field.
func (r *videoResolver) Status(ctx context.Context, obj *model.Video) (*model.Status, error) {
	return r.VideoService.GetStatus(obj.ID)
}

// Profile is the resolver for the profile field.
func (r *videoResolver) Profile(ctx context.Context, obj *model.Video) (*model.Profile, error) {
	return r.VideoService.GetProfile(obj.ProfileID)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Video returns VideoResolver implementation.
func (r *Resolver) Video() VideoResolver { return &videoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type videoResolver struct{ *Resolver }
