package user

import "time"

type User struct {
	ID int64
	Name string
	Password string
	salt string
	admin bool
	expiredAt time.Time
}
