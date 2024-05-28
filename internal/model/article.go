package model

import "time"

type Article struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    string    `json:"-"`
	Votes     int       `json:"votes"`
	IsClosed  bool      `json:"isClosed"`
	CreatedAt time.Time `json:"createdAt"`
}
