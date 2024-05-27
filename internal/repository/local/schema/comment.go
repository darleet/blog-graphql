package schema

import "time"

type Comment struct {
	ID        uint64
	Content   string
	UserID    uint64
	Parent    uint64
	Votes     map[uint64]int
	CreatedAt time.Time
}
