package comment

import "time"

type Comment struct {
	ID int64
	Comment string
	CommentedAt time.Time
}
