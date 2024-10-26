package dto

import "github.com/ValkyrieKia/golang-deals-test-project/internal/app/entity"

type ItemRequestDTO struct {
	Name           string `json:"name" binding:"required"`
	IdUnit         int32  `json:"id_unit" binding:"required"`
	IdItemCategory int32  `json:"id_item_category" binding:"required"`
}

func (h *ItemRequestDTO) ToEntity() *entity.Item {
	if h == nil {
		return nil
	}

	return &entity.Item{
		Name:           h.Name,
		IdUnit:         h.IdUnit,
		IdItemCategory: h.IdItemCategory,
	}

}
