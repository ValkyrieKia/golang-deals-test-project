package entity

import "time"

type User struct {
	ID        int
	Username  string
	FullName  string
	Password  string
	CreatedAt time.Time
}
