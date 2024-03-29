// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CommentInput struct {
	Body    string `json:"body"`
	VideoID string `json:"VideoId"`
}

type ContentInformation struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Published   string `json:"published"`
	Channelid   string `json:"channelid"`
}

type Profile struct {
	Username           string  `json:"username"`
	Displayname        *string `json:"displayname"`
	IsChannel          *bool   `json:"isChannel"`
	Subscribers        *int    `json:"subscribers"`
	UserIsSubscribedTo *bool   `json:"userIsSubscribedTo"`
}

type Statistic struct {
	Likes     int  `json:"likes"`
	Dislikes  int  `json:"dislikes"`
	Views     int  `json:"views"`
	Favorites *int `json:"favorites"`
	Comments  int  `json:"comments"`
}

type Status struct {
	Uploadstatus  Uploadstatus  `json:"uploadstatus"`
	Privacystatus Privacystatus `json:"privacystatus"`
}

type Thumbnail struct {
	URL string `json:"url"`
}

type VideoInput struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Username    string `json:"username"`
	Displayname string `json:"displayname"`
	Password    string `json:"password"`
	Password2   string `json:"password2"`
}

type LikeDislike string

const (
	LikeDislikeLike    LikeDislike = "like"
	LikeDislikeDislike LikeDislike = "dislike"
)

var AllLikeDislike = []LikeDislike{
	LikeDislikeLike,
	LikeDislikeDislike,
}

func (e LikeDislike) IsValid() bool {
	switch e {
	case LikeDislikeLike, LikeDislikeDislike:
		return true
	}
	return false
}

func (e LikeDislike) String() string {
	return string(e)
}

func (e *LikeDislike) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LikeDislike(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LIKE_DISLIKE", str)
	}
	return nil
}

func (e LikeDislike) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Privacystatus string

const (
	PrivacystatusPrivate Privacystatus = "private"
	PrivacystatusPublic  Privacystatus = "public"
)

var AllPrivacystatus = []Privacystatus{
	PrivacystatusPrivate,
	PrivacystatusPublic,
}

func (e Privacystatus) IsValid() bool {
	switch e {
	case PrivacystatusPrivate, PrivacystatusPublic:
		return true
	}
	return false
}

func (e Privacystatus) String() string {
	return string(e)
}

func (e *Privacystatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Privacystatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PRIVACYSTATUS", str)
	}
	return nil
}

func (e Privacystatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Uploadstatus string

const (
	UploadstatusProcessing Uploadstatus = "processing"
	UploadstatusError      Uploadstatus = "error"
	UploadstatusComplete   Uploadstatus = "complete"
)

var AllUploadstatus = []Uploadstatus{
	UploadstatusProcessing,
	UploadstatusError,
	UploadstatusComplete,
}

func (e Uploadstatus) IsValid() bool {
	switch e {
	case UploadstatusProcessing, UploadstatusError, UploadstatusComplete:
		return true
	}
	return false
}

func (e Uploadstatus) String() string {
	return string(e)
}

func (e *Uploadstatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Uploadstatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UPLOADSTATUS", str)
	}
	return nil
}

func (e Uploadstatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
