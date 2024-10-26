package repository

import "github.com/ValkyrieKia/golang-deals-test-project/internal/app/entity"

type UserSessionRepository interface {
	GetByUid(sessionUid string) (*entity.UserSession, error)
	Create(session *entity.UserSession) (*entity.UserSession, error)
	Update(uid string, session *entity.UserSession) (*entity.UserSession, error)
	Destroy(sessionUid string) error
}
