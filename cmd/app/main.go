package main

import (
	"log"

	"github.com/ValkyrieKia/golang-deals-test-project/internal/config"
	rest "github.com/ValkyrieKia/golang-deals-test-project/internal/controller"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/controller/common"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/infrastructure/provider"
)

func main() {
	cfg := config.GetConfig()

	dbMySql, err := provider.InitDbConn("mysql", &cfg.DatabaseConfig, cfg.GenericConfig.AppEnv == "development")

	if err != nil {
		log.Fatalf("Database connection failed: %s", err)
	}

	rest.InitHttpServer(&common.HttpServerProviders{
		Config:           cfg,
		MainDbConnection: dbMySql,
	})
}
