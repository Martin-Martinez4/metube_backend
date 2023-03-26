package testhelpers

import "github/Martin-Martinez4/metube_backend/graph/model"

var coding_channel = "coding"
var Coding_channel_id = "2d089bcf-1813-4f99-bad1-bd5640dcf0bc"
var TRUE = true
var codingSubscribers = 3300

var CodingChannelProfile = &model.Profile{
	Username:    "coding_channel",
	Displayname: &coding_channel,
	IsChannel:   &TRUE,
	Subscribers: &codingSubscribers,
}

var FALSE = false
var test_channel = "test_channel"
var testSubscribers = 0

var TestChannelProfile = &model.Profile{
	Username:    "test_channel",
	Displayname: &test_channel,
	IsChannel:   &FALSE,
	Subscribers: &testSubscribers,
}

var anime_channel = "anime_channel"
var Anime_channel_id = "6adbb5ec-13c3-46c8-9b94-3c9f2cf1a660"
var animeSubscribers = 200000

var AnimeChannelProfile = &model.Profile{
	Username:    "anime_chanbel",
	Displayname: &anime_channel,
	IsChannel:   &TRUE,
	Subscribers: &animeSubscribers,
}
