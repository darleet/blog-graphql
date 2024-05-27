package model

import "time"

type Comment struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	UserID    string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}
