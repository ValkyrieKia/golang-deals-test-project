package item

import (
	"context"

	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/entity"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/repository"
	"gorm.io/gorm"
)

type itemMysqlRepository struct {
	db *gorm.DB
}

func NewitemMysqlRepository(
	db *gorm.DB,
) repository.ItemRepository {
	return &itemMysqlRepository{
		db: db,
	}
}

func (u itemMysqlRepository) GetAll(ctx context.Context) ([]*entity.Item, error) {
	var queryResult []ItemCommonModel

	stmt := u.db.Model(ItemCommonModel{}).Select("*").Where("visible = ?", 1)
	tx := stmt.Find(&queryResult)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return ListItemCommonModel(queryResult).ToEntities(), nil
}

func (u itemMysqlRepository) Create(ctx context.Context, data entity.Item) error {
	model := ItemCommonModel{}.FromEntity(&data)
	tx := u.db.Create(&model)
	err := tx.Error

	if err != nil {
		return err
	}

	return nil
}
