package user_session

import (
	"time"

	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/entity"
)

type UserSessionModel struct {
	ID           int       `gorm:"column:id;primary_key"`
	UID          string    `gorm:"column:uid"`
	UserID       int       `gorm:"column:user_id"`
	RefreshToken string    `gorm:"column:refresh_token"`
	DeviceInfo   string    `gorm:"column:device_info;default:NULL"`
	IPAddress    string    `gorm:"column:ip_address;default:NULL"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	ExpiresAt    time.Time `gorm:"column:expires_at"`
}

func (m UserSessionModel) TableName() string {
	return "user_session"
}

func (m UserSessionModel) ToEntity() *entity.UserSession {
	sess := &entity.UserSession{
		SessionUID:   m.UID,
		User:         &entity.User{ID: m.UserID},
		RefreshToken: m.RefreshToken,
		IPAddress:    m.IPAddress,
		ExpiresAt:    m.ExpiresAt,
		CreatedAt:    m.CreatedAt,
	}
	if m.DeviceInfo != "" {
		sess.DeviceInfo = &m.DeviceInfo
	}
	return sess
}

func (m UserSessionModel) FromEntity(e *entity.UserSession) *UserSessionModel {
	res := &UserSessionModel{
		UID:          e.SessionUID,
		RefreshToken: e.RefreshToken,
		IPAddress:    e.IPAddress,
		CreatedAt:    e.CreatedAt,
		ExpiresAt:    e.ExpiresAt,
	}

	if e.User != nil {
		res.UserID = e.User.ID
	}
	if e.DeviceInfo != nil {
		res.DeviceInfo = *e.DeviceInfo
	}

	return res
}
