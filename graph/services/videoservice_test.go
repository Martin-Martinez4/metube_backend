package services

import (
	"github/Martin-Martinez4/metube_backend/graph/model"
	helpers "github/Martin-Martinez4/metube_backend/graph/services/testHelpers"
	"reflect"
	"testing"
)

var armCodingVideoId = "ea232a13-eb2d-429f-ab22-579ffaf5d6b0"

func TestVideoServiceSQL_GetVideoById(t *testing.T) {

	DB := helpers.StartTestDB()
	defer DB.Close()

	videoService := &VideoServiceSQL{DB: DB}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		vsql    *VideoServiceSQL
		args    args
		want    *model.Video
		wantErr bool
	}{
		{
			name: "get video by id",
			vsql: videoService,
			args: args{
				id: armCodingVideoId,
			},
			want:    &helpers.ARMCodingVideo,
			wantErr: false,
		},
		{
			name: "empty id",
			vsql: videoService,
			args: args{
				id: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.vsql.GetVideoById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("VideoServiceSQL.GetVideoById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VideoServiceSQL.GetVideoById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVideoServiceSQL_GetMultipleVideos(t *testing.T) {
	type args struct {
		amount int
	}
	tests := []struct {
		name    string
		vsql    *VideoServiceSQL
		args    args
		want    []*model.Video
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.vsql.GetMultipleVideos(tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("VideoServiceSQL.GetMultipleVideos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VideoServiceSQL.GetMultipleVideos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVideoServiceSQL_GetContentInformation(t *testing.T) {

	DB := helpers.StartTestDB()
	defer DB.Close()

	videoService := &VideoServiceSQL{DB: DB}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		vsql    *VideoServiceSQL
		args    args
		want    *model.ContentInformation
		wantErr bool
	}{
		{
			name: "get content information",
			vsql: videoService,
			args: args{
				id: armCodingVideoId,
			},
			want:    &helpers.ARMCodingVideoContentInformation,
			wantErr: false,
		},
		{
			name: "empty id",
			vsql: videoService,
			args: args{
				id: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.vsql.GetContentInformation(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("VideoServiceSQL.GetContentInformation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VideoServiceSQL.GetContentInformation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVideoServiceSQL_GetThumbnail(t *testing.T) {

	DB := helpers.StartTestDB()
	defer DB.Close()

	videoService := &VideoServiceSQL{DB: DB}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		vsql    *VideoServiceSQL
		args    args
		want    *model.Thumbnail
		wantErr bool
	}{
		{
			name: "get video thumbnail",
			vsql: videoService,
			args: args{
				id: armCodingVideoId,
			},
			want:    &helpers.ARMCodingVideoThumbnail,
			wantErr: false,
		},
		{
			name: "empty id",
			vsql: videoService,
			args: args{
				id: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.vsql.GetThumbnail(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("VideoServiceSQL.GetThumbnail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VideoServiceSQL.GetThumbnail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVideoServiceSQL_GetStatistic(t *testing.T) {

	DB := helpers.StartTestDB()
	defer DB.Close()

	videoService := &VideoServiceSQL{DB: DB}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		vsql    *VideoServiceSQL
		args    args
		want    *model.Statistic
		wantErr bool
	}{
		{
			name: "get video statistics",
			vsql: videoService,
			args: args{
				id: armCodingVideoId,
			},
			want:    &helpers.ARMCodingVideoStatistic,
			wantErr: false,
		},
		{
			name: "empty id",
			vsql: videoService,
			args: args{
				id: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.vsql.GetStatistic(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("VideoServiceSQL.GetStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VideoServiceSQL.GetStatistic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVideoServiceSQL_GetStatus(t *testing.T) {

	DB := helpers.StartTestDB()
	defer DB.Close()

	videoService := &VideoServiceSQL{DB: DB}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		vsql    *VideoServiceSQL
		args    args
		want    *model.Status
		wantErr bool
	}{
		{
			name: "get video statistics",
			vsql: videoService,
			args: args{
				id: armCodingVideoId,
			},
			want:    &helpers.ARMCodingVideoStatus,
			wantErr: false,
		},
		{
			name: "empty id",
			vsql: videoService,
			args: args{
				id: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.vsql.GetStatus(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("VideoServiceSQL.GetStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VideoServiceSQL.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestVideoServiceSQL_GetProfile(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		vsql    *VideoServiceSQL
		args    args
		want    *model.Profile
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.vsql.GetProfile(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("VideoServiceSQL.GetProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VideoServiceSQL.GetProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
