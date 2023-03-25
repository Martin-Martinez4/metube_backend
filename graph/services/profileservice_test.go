package services

import (
	"context"
	"github/Martin-Martinez4/metube_backend/graph/model"
	helpers "github/Martin-Martinez4/metube_backend/graph/services/testHelpers"
	"github/Martin-Martinez4/metube_backend/utils"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"testing"
)

func TestProfileServiceSQL_GetProfileIdFromUsername(t *testing.T) {

	DB := helpers.StartTestDB()
	defer DB.Close()

	type args struct {
		username string
	}
	tests := []struct {
		name    string
		psql    *ProfileServiceSQL
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Get id from username",
			psql: &ProfileServiceSQL{DB: DB},
			args: args{
				username: "coding_channel",
			},
			want:    "2d089bcf-1813-4f99-bad1-bd5640dcf0bc",
			wantErr: false,
		},
		{
			name: "invalid username",
			psql: &ProfileServiceSQL{DB: DB},
			args: args{
				username: "invalid",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.psql.GetProfileIdFromUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileServiceSQL.GetProfileIdFromUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ProfileServiceSQL.GetProfileIdFromUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileServiceSQL_GetProfileByUsername(t *testing.T) {

	DB := helpers.StartTestDB()
	defer DB.Close()

	type args struct {
		username string
	}
	tests := []struct {
		name    string
		psql    *ProfileServiceSQL
		args    args
		want    *model.Profile
		wantErr bool
	}{
		{
			name: "successful getting profile",
			psql: &ProfileServiceSQL{DB: DB},
			args: args{
				username: "coding_channel",
			},
			want:    helpers.CodingChannelProfile,
			wantErr: false,
		},
		{
			name: "failure getting profile",
			psql: &ProfileServiceSQL{DB: DB},
			args: args{
				username: "AMistake;k",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.psql.GetProfileByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileServiceSQL.GetProfileByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileServiceSQL.GetProfileByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileServiceSQL_GetMultipleProfiles(t *testing.T) {

	DB := helpers.StartTestDB()
	defer DB.Close()

	type args struct {
		amount int
	}
	tests := []struct {
		name    string
		psql    *ProfileServiceSQL
		args    args
		want    []*model.Profile
		wantErr bool
	}{
		{
			name: "get multiple profiles success",
			psql: &ProfileServiceSQL{DB: DB},
			args: args{
				amount: 3,
			},
			want: []*model.Profile{
				helpers.CodingChannelProfile,
				helpers.AnimeChannelProfile,
				helpers.TestChannelProfile,
			},
			wantErr: false,
		},
		{
			name: "amount bigger than available profiles",
			psql: &ProfileServiceSQL{DB: DB},
			args: args{
				amount: 10,
			},
			want: []*model.Profile{
				helpers.CodingChannelProfile,
				helpers.AnimeChannelProfile,
				helpers.TestChannelProfile,
			},
			wantErr: false,
		},
		{
			name: "get multiple profiles get zero profiles",
			psql: &ProfileServiceSQL{DB: DB},
			args: args{
				amount: 0,
			},
			want:    []*model.Profile{},
			wantErr: false,
		},
		{
			name: "get multiple profiles get zero profiles",
			psql: &ProfileServiceSQL{DB: DB},
			args: args{
				amount: -1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.psql.GetMultipleProfiles(tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileServiceSQL.GetMultipleProfiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			sort.SliceStable(got, func(i, j int) bool {
				return *got[i].Displayname < *got[j].Displayname
			})

			sort.SliceStable(tt.want, func(i, j int) bool {
				return *tt.want[i].Displayname < *tt.want[j].Displayname
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileServiceSQL.GetMultipleProfiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileServiceSQL_Subscribe(t *testing.T) {

	DB := helpers.StartTestDB()
	defer DB.Close()
	subscriberId := "6adbb5ec-13c3-46c8-9b94-3c9f2cf1a660"

	type args struct {
		subscribee string
	}
	tests := []struct {
		name          string
		psql          *ProfileServiceSQL
		args          args
		subscriber_id string
		want          bool
		wantErr       bool
	}{
		{
			name: "subscribe to a channel",
			psql: &ProfileServiceSQL{DB: DB},
			args: args{

				subscribee: "coding_channel",
			},
			subscriber_id: subscriberId,
			want:          true,
			wantErr:       false,
		},
		{
			name: "duplicate subscription, return true but nothing happens in the DB",
			psql: &ProfileServiceSQL{DB: DB},
			args: args{

				subscribee: "coding_channel",
			},
			subscriber_id: subscriberId,
			want:          true,
			wantErr:       false,
		},
		{
			name: "subscribe to oneself, is allowed",
			psql: &ProfileServiceSQL{DB: DB},
			args: args{

				subscribee: "anime_chanbel",
			},
			subscriber_id: subscriberId,
			want:          true,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Mock the request to place the user id in the request's context
			req := httptest.NewRequest(http.MethodPost, "/query", nil)
			ctx2 := context.WithValue(req.Context(), utils.UserKey, tt.subscriber_id)

			got, err := tt.psql.Subscribe(ctx2, tt.args.subscribee)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileServiceSQL.Subscribe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ProfileServiceSQL.Subscribe() = %v, want %v", got, tt.want)
			}
		})
	}
}
