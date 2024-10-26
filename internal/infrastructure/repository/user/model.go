package user

import (
	"time"

	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/entity"
)

type UserModel struct {
	ID        int       `gorm:"column:id;primaryKey"`
	Username  string    `gorm:"column:username"`
	FullName  string    `gorm:"column:full_name"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at;default:current_timestamp"`
}

func (u UserModel) TableName() string {
	return "user"
}

func (u UserModel) ToEntity() *entity.User {
	return &entity.User{
		ID:        u.ID,
		Username:  u.Username,
		FullName:  u.FullName,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
	}
}
