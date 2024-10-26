package repository

import "github.com/ValkyrieKia/golang-deals-test-project/internal/app/entity"

type UserRepository interface {
	GetByUsername(username string) (*entity.User, error)
}
