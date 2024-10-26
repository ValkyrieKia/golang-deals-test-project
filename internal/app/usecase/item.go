package usecase

import (
	"context"

	"github.com/ValkyrieKia/golang-deals-test-project/internal/util"

	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/entity"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/repository"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/config"
)

type ItemUsecase struct {
	config         *config.Config
	itemRepository repository.ItemRepository
}

func NewItemUsecase(
	config *config.Config,
	itemRepository repository.ItemRepository,
) *ItemUsecase {
	return &ItemUsecase{
		config:         config,
		itemRepository: itemRepository,
	}
}

func (au *ItemUsecase) GetAll(ctx context.Context) ([]*entity.Item, error) {
	item, err := au.itemRepository.GetAll(ctx)
	if err != nil {
		return nil, util.NewCommonError(err, util.ErrInternal, err.Error())
	}

	return item, nil
}

func (au *ItemUsecase) Create(ctx context.Context, data entity.Item) error {
	err := au.itemRepository.Create(ctx, data)
	if err != nil {
		return util.NewCommonError(err, util.ErrInternal, err.Error())
	}
	return nil
}
