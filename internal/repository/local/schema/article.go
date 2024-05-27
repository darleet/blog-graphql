package schema

import "time"

type Article struct {
	ID        uint64
	Title     string
	Content   string
	UserID    uint64
	Comments  []uint64
	IsClosed  bool
	Votes     map[uint64]int
	CreatedAt time.Time
}
