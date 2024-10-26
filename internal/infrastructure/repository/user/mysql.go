package user

import (
	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/entity"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/repository"
	"gorm.io/gorm"
)

type userMysqlRepository struct {
	db *gorm.DB
}

func NewUserMysqlRepository(
	db *gorm.DB,
) repository.UserRepository {
	return &userMysqlRepository{
		db: db,
	}
}

func (u userMysqlRepository) GetByUsername(username string) (*entity.User, error) {
	stmt := u.db.Model(UserModel{}).Select("*").Where("username = ?", username)

	var queryResult UserModel
	tx := stmt.First(&queryResult)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return queryResult.ToEntity(), nil
}
