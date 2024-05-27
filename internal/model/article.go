package model

import "time"

type Article struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    string    `json:"-"`
	IsClosed  bool      `json:"isClosed"`
	Votes     int       `json:"votes"`
	CreatedAt time.Time `json:"createdAt"`
}
