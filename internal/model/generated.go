// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type LoginInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Mutation struct {
}

type NewArticle struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	IsClosed bool   `json:"isClosed"`
}

type NewComment struct {
	ArticleID string  `json:"articleID"`
	Content   string  `json:"content"`
	ParentID  *string `json:"parentID,omitempty"`
}

type Query struct {
}

type RegisterInput struct {
	Username  string `json:"username"`
	Email     Email  `json:"email"`
	Password  string `json:"password"`
	AvatarURL *URL   `json:"avatarURL,omitempty"`
}

type Subscription struct {
}

type UpdateArticle struct {
	ID       string  `json:"id"`
	Title    *string `json:"title,omitempty"`
	Content  *string `json:"content,omitempty"`
	IsClosed *bool   `json:"isClosed,omitempty"`
}

type UpdateComment struct {
	ID      string  `json:"id"`
	Content *string `json:"content,omitempty"`
}

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	AvatarURL *URL   `json:"avatarURL,omitempty"`
}

type Vote struct {
	ArticleID string    `json:"articleID"`
	Value     VoteValue `json:"value"`
}

type Sort string

const (
	SortTopAsc  Sort = "TOP_ASC"
	SortTopDesc Sort = "TOP_DESC"
	SortNewAsc  Sort = "NEW_ASC"
	SortNewDesc Sort = "NEW_DESC"
)

var AllSort = []Sort{
	SortTopAsc,
	SortTopDesc,
	SortNewAsc,
	SortNewDesc,
}

func (e Sort) IsValid() bool {
	switch e {
	case SortTopAsc, SortTopDesc, SortNewAsc, SortNewDesc:
		return true
	}
	return false
}

func (e Sort) String() string {
	return string(e)
}

func (e *Sort) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Sort(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Sort", str)
	}
	return nil
}

func (e Sort) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type VoteValue string

const (
	VoteValueNone VoteValue = "NONE"
	VoteValueUp   VoteValue = "UP"
	VoteValueDown VoteValue = "DOWN"
)

var AllVoteValue = []VoteValue{
	VoteValueNone,
	VoteValueUp,
	VoteValueDown,
}

func (e VoteValue) IsValid() bool {
	switch e {
	case VoteValueNone, VoteValueUp, VoteValueDown:
		return true
	}
	return false
}

func (e VoteValue) String() string {
	return string(e)
}

func (e *VoteValue) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VoteValue(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VoteValue", str)
	}
	return nil
}

func (e VoteValue) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
