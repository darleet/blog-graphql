package model

import "time"

type Comment struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	UserID    string    `json:"-"`
	Votes     int       `json:"votes"`
	CreatedAt time.Time `json:"createdAt"`
}
