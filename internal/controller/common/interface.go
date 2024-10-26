package common

import (
	"github.com/ValkyrieKia/golang-deals-test-project/internal/config"
	"gorm.io/gorm"
)

type HttpServerProviders struct {
	Config           *config.Config
	MainDbConnection *gorm.DB
}
