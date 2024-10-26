package entity

import "time"

type UserSession struct {
	SessionUID   string
	User         *User
	RefreshToken string
	IPAddress    string
	DeviceInfo   *string
	ExpiresAt    time.Time
	CreatedAt    time.Time
}
