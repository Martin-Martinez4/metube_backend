package model

type Video struct {
	ID                 string              `json:"id"`
	URL                string              `json:"url"`
	Categoryid         string              `json:"categoryid"`
	Duration           int                 `json:"duration"`
	ProfileID          string              `json:"profile_id"`
	Contentinformation *ContentInformation `json:"contentinformation"`
	Thumbnail          *Thumbnail          `json:"thumbnail"`
	Statistic          *Statistic          `json:"statistic"`
	Status             *Status             `json:"status"`
	Profile            *Profile            `json:"profile"`
}
