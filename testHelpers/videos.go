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

var SQLCodingVideoContentInformation = model.ContentInformation{
	Title:       "SQL Server Performance Essentials",
	Description: "The Essentials of SQL server performance.",
	Published:   "2023-03-14T19:39:22+00:00",
	Channelid:   "2d089bcf-1813-4f99-bad1-bd5640dcf0bc",
}

var SQLCodingVideoThumbnail = model.Thumbnail{
	URL: "/thumbnail/SQL_Server_Performance_Essentials.jpg/",
}

var SQLCodingVideoStatistic = model.Statistic{
	Likes:     20,
	Dislikes:  0,
	Views:     200,
	Comments:  20,
	Favorites: &favorites,
}

var SQLCodingVideoStatus = model.Status{
	Uploadstatus:  "complete",
	Privacystatus: "public",
}

var SQLCodingVideo = model.Video{
	ID:                 "3d7eeca5-f0d5-4752-b213-b8e9445627f2",
	URL:                "/media/SQL_Server_Performance_Essentials/stream/",
	Categoryid:         "coding",
	Duration:           14606,
	ProfileID:          "2d089bcf-1813-4f99-bad1-bd5640dcf0bc",
	Contentinformation: nil,
	Thumbnail:          nil,
	Statistic:          nil,
	Status:             nil,
	Profile:            nil,
}

var JapaneseCartoonVideoContentInformation = model.ContentInformation{
	Title:       "Trigun Episode 1",
	Description: "The first episode of Trigun",
	Published:   "2023-03-14T19:39:22+00:00",
	Channelid:   "6adbb5ec-13c3-46c8-9b94-3c9f2cf1a660",
}

var JapaneseCartoonVideoThumbnail = model.Thumbnail{
	URL: "/thumbnail/Trigun_Ep1.jpg/",
}

var JapaneseCartoonVideoStatistic = model.Statistic{
	Likes:     20,
	Dislikes:  0,
	Views:     200,
	Comments:  20,
	Favorites: &favorites,
}

var JapaneseCartoonVideoStatus = model.Status{
	Uploadstatus:  "complete",
	Privacystatus: "public",
}

var JapaneseCartoonVideo = model.Video{
	ID:                 "75f665c6-0aa4-4463-bb57-854c66c9bed8",
	URL:                "/media/Trigun_Ep1/stream/",
	Categoryid:         "anime",
	Duration:           1440,
	ProfileID:          "6adbb5ec-13c3-46c8-9b94-3c9f2cf1a660",
	Contentinformation: nil,
	Thumbnail:          nil,
	Statistic:          nil,
	Status:             nil,
	Profile:            nil,
}
