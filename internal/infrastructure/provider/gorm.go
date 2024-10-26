package provider

import (
	"errors"
	"fmt"

	"github.com/ValkyrieKia/golang-deals-test-project/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type DBOption func(db *gorm.DB) *gorm.DB

type DBConnectionProps struct {
	Hostname     string
	ReadHostname string
	Username     string
	Password     string
	DBName       string
	DebugMode    bool
}

func InitDbConn(driverName string, dbConfig *config.DatabaseConfig, debug bool) (*gorm.DB, error) {
	coreDsn := generateDsnString(
		driverName, dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Database,
	)
	replicaDsn := generateDsnString(
		driverName, dbConfig.Username, dbConfig.Password, dbConfig.HostRead, dbConfig.Database,
	)

	gormOption := &gorm.Config{}

	if debug == true {
		gormOption = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	var db *gorm.DB
	var err error
	switch driverName {
	case "mysql":
		db, err = gorm.Open(mysql.Open(coreDsn), gormOption)
		if db != nil {
			err = db.Use(dbresolver.Register(dbresolver.Config{
				Replicas:          []gorm.Dialector{mysql.Open(replicaDsn)},
				TraceResolverMode: debug,
			}))

		}
	default:
		return nil, errors.New("unimplemented driver")
	}
	return db, err
}

func generateDsnString(driver, username, password, hostname, dbName string) string {
	switch driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, dbName)
	}
	return ""
}
