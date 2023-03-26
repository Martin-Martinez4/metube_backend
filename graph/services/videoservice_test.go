package services

import (
	config "github/Martin-Martinez4/metube_backend/config"
	"github/Martin-Martinez4/metube_backend/graph/model"
	helpers "github/Martin-Martinez4/metube_backend/testHelpers"
	"reflect"
	"sort"
	"testing"
)

var armCodingVideoId = "ea232a13-eb2d-429f-ab22-579ffaf5d6b0"
var SQLCodingVideo = "3d7eeca5-f0d5-4752-b213-b8e9445627f2"

func TestVideoServiceSQL_GetVideoById(t *testing.T) {

	TEST_DB_URL := config.ReadEnv("../../.env").TEST_DB_URL

	DB := config.GetDB("postgres", TEST_DB_URL)
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
			name: "get arm coding video by id",
			vsql: videoService,
			args: args{
				id: armCodingVideoId,
			},
			want:    &helpers.ARMCodingVideo,
			wantErr: false,
		},
		{
			name: "get sql coding video by id",
			vsql: videoService,
			args: args{
				id: SQLCodingVideo,
			},
			want:    &helpers.SQLCodingVideo,
			wantErr: false,
		},
		{
			name: "malformed id",
			vsql: &VideoServiceSQL{DB: DB},
			args: args{
				id: "a465sd-465465465",
			},
			want:    nil,
			wantErr: true,
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
	TEST_DB_URL := config.ReadEnv("../../.env").TEST_DB_URL

	DB := config.GetDB("postgres", TEST_DB_URL)
	defer DB.Close()

	videoService := &VideoServiceSQL{DB: DB}

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
		{
			name: "get video 10 random videos",
			vsql: videoService,
			args: args{
				amount: 10,
			},
			want:    []*model.Video{&helpers.ARMCodingVideo, &helpers.SQLCodingVideo, &helpers.JapaneseCartoonVideo},
			wantErr: false,
		},
		{
			name: "get video 0 videos",
			vsql: videoService,
			args: args{
				amount: 0,
			},
			want:    []*model.Video{},
			wantErr: false,
		},
		{
			name: "invaild value for amount",
			vsql: videoService,
			args: args{
				amount: -1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.vsql.GetMultipleVideos(tt.args.amount)

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].URL < got[j].URL
			})

			sort.SliceStable(tt.want, func(i, j int) bool {
				return tt.want[i].URL < tt.want[j].URL
			})
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

	TEST_DB_URL := config.ReadEnv("../../.env").TEST_DB_URL

	DB := config.GetDB("postgres", TEST_DB_URL)
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
			name: "malformed id",
			vsql: &VideoServiceSQL{DB: DB},
			args: args{
				id: "a465sd-465465465",
			},
			want:    nil,
			wantErr: true,
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

	TEST_DB_URL := config.ReadEnv("../../.env").TEST_DB_URL

	DB := config.GetDB("postgres", TEST_DB_URL)
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
			name: "malformed id",
			vsql: &VideoServiceSQL{DB: DB},
			args: args{
				id: "a465sd-465465465",
			},
			want:    nil,
			wantErr: true,
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

	TEST_DB_URL := config.ReadEnv("../../.env").TEST_DB_URL

	DB := config.GetDB("postgres", TEST_DB_URL)
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
			name: "malformed id",
			vsql: &VideoServiceSQL{DB: DB},
			args: args{
				id: "a465sd-465465465",
			},
			want:    nil,
			wantErr: true,
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

	TEST_DB_URL := config.ReadEnv("../../.env").TEST_DB_URL

	DB := config.GetDB("postgres", TEST_DB_URL)
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
			name: "malformed id",
			vsql: &VideoServiceSQL{DB: DB},
			args: args{
				id: "a465sd-465465465",
			},
			want:    nil,
			wantErr: true,
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

	TEST_DB_URL := config.ReadEnv("../../.env").TEST_DB_URL

	DB := config.GetDB("postgres", TEST_DB_URL)
	defer DB.Close()

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
		{
			name: "get coding channel",
			vsql: &VideoServiceSQL{DB: DB},
			args: args{
				id: helpers.Coding_channel_id,
			},
			want:    helpers.CodingChannelProfile,
			wantErr: false,
		},
		{
			name: "malformed id",
			vsql: &VideoServiceSQL{DB: DB},
			args: args{
				id: "a465sd-465465465",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty id",
			vsql: &VideoServiceSQL{DB: DB},
			args: args{
				id: "",
			},
			want:    nil,
			wantErr: true,
		},
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
