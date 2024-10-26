package item

import (
	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/entity"
)

type ItemCommonModel struct {
	ID             int32  `gorm:"column:id_item"`
	Name           string `gorm:"column:name"`
	IdUnit         int32  `gorm:"column:id_unit"`
	IdItemCategory int32  `gorm:"column:id_item_category"`
	Visible        int32  `gorm:"column:visible"`
}

func (u ItemCommonModel) TableName() string {
	return "item"
}

func (u ItemCommonModel) ToEntity() *entity.Item {
	return &entity.Item{
		ID:             u.ID,
		Name:           u.Name,
		IdUnit:         u.IdUnit,
		IdItemCategory: u.IdItemCategory,
		Visible:        u.Visible,
	}
}

func (m ItemCommonModel) FromEntity(ent *entity.Item) ItemCommonModel {
	return ItemCommonModel{
		ID:             m.ID,
		Name:           m.Name,
		IdUnit:         m.IdUnit,
		IdItemCategory: m.IdItemCategory,
		Visible:        m.Visible,
	}
}

type ListItemCommonModel []ItemCommonModel

func (model ListItemCommonModel) ToEntities() []*entity.Item {
	var entities []*entity.Item

	for _, data := range model {
		entities = append(entities, data.ToEntity())
	}

	return entities
}
