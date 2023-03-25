package testhelpers

import "github/Martin-Martinez4/metube_backend/graph/model"

var ARMCodingVideoContentInformation = model.ContentInformation{
	Title:       "Assembly Language Programming with ARM",
	Description: "Learn Asembly with ARM",
	Published:   "2023-03-14T19:39:22+00:00",
	Channelid:   "2d089bcf-1813-4f99-bad1-bd5640dcf0bc",
}

var ARMCodingVideoThumbnail = model.Thumbnail{
	URL: "/thumbnail/Assembly_Language_Programming_with_ARM.jpg/",
}

var favorites = 2

var ARMCodingVideoStatistic = model.Statistic{
	Likes:     20,
	Dislikes:  0,
	Views:     200,
	Comments:  20,
	Favorites: &favorites,
}

var ARMCodingVideoStatus = model.Status{
	Uploadstatus:  "complete",
	Privacystatus: "public",
}

var ARMCodingVideo = model.Video{
	ID:                 "ea232a13-eb2d-429f-ab22-579ffaf5d6b0",
	URL:                "/media/Assembly_Language_Programming_with_ARM/stream/",
	Categoryid:         "coding",
	Duration:           8971,
	ProfileID:          "2d089bcf-1813-4f99-bad1-bd5640dcf0bc",
	Contentinformation: nil,
	Thumbnail:          nil,
	Statistic:          nil,
	Status:             nil,
	Profile:            nil,
}
