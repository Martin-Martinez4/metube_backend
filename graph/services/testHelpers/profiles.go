package testhelpers

import "github/Martin-Martinez4/metube_backend/graph/model"

var coding_channel = "coding"
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
var animeSubscribers = 200000

var AnimeChannelProfile = &model.Profile{
	Username:    "anime_chanbel",
	Displayname: &anime_channel,
	IsChannel:   &TRUE,
	Subscribers: &animeSubscribers,
}
