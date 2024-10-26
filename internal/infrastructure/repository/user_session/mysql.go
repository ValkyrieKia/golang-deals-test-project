package user_session

import (
	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/entity"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/repository"
	"gorm.io/gorm"
)

type userSessionMysqlRepository struct {
	db *gorm.DB
}

func NewUserSessionMysqlRepository(
	db *gorm.DB,
) repository.UserSessionRepository {
	return &userSessionMysqlRepository{
		db: db,
	}
}

func (u userSessionMysqlRepository) GetByUid(sessionUid string) (*entity.UserSession, error) {
	stmt := u.db.Model(UserSessionModel{}).Where("uid = ?", sessionUid)

	var queryResult *UserSessionModel
	tx := stmt.First(&queryResult)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return queryResult.ToEntity(), nil
}

func (u userSessionMysqlRepository) Create(session *entity.UserSession) (*entity.UserSession, error) {
	model := UserSessionModel{}.FromEntity(session)
	trx := u.db.Create(model)
	if trx.Error != nil {
		return nil, trx.Error
	}
	return model.ToEntity(), nil
}

func (u userSessionMysqlRepository) Update(uid string, session *entity.UserSession) (*entity.UserSession, error) {
	model := UserSessionModel{}.FromEntity(session)
	stmt := u.db.Model(UserSessionModel{}).Where("uid = ?", uid)
	trx := stmt.Updates(model)
	if trx.Error != nil {
		return nil, trx.Error
	}
	return model.ToEntity(), nil
}

func (u userSessionMysqlRepository) Destroy(sessionUid string) error {
	stmt := u.db.Model(UserSessionModel{}).Where("uid = ?", sessionUid)
	trx := stmt.Delete(&UserSessionModel{})
	if trx.Error != nil {
		return trx.Error
	}
	return nil
}
