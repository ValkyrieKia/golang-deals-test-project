package repository

import (
	"context"

	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/entity"
)

type ItemRepository interface {
	GetAll(ctx context.Context) ([]*entity.Item, error)
	Create(ctx context.Context, data entity.Item) error
}
