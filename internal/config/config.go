package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	GenericConfig  GenericConfig
	DatabaseConfig DatabaseConfig
	AuthConfig     AuthConfig
}

type GenericConfig struct {
	AppEnv  string `envconfig:"APP_ENV" default:"development"`
	AppPort string `envconfig:"APP_PORT" default:"8080"`
}

type DatabaseConfig struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	HostRead string `envconfig:"DB_HOST_READ" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"3306"`
	Username string `envconfig:"DB_USERNAME" default:"root"`
	Password string `envconfig:"DB_PASSWORD" default:"password"`
	Database string `envconfig:"DB_DATABASE" default:"mysql"`
}

type AuthConfig struct {
	JwtTokenSecret        string `envconfig:"JWT_TOKEN_SECRET" default:"secret"`
	JwtRefreshTokenSecret string `envconfig:"JWT_REFRESH_TOKEN_SECRET" default:"secret"`
	JwtTokenExpiry        string `envconfig:"JWT_TOKEN_EXPIRY" default:"1h"`
	JwtRefreshTokenExpiry string `envconfig:"JWT_REFRESH_TOKEN_EXPIRY" default:"720h"`
}

var cfg Config

func GetConfig() *Config {
	envconfig.MustProcess("", &cfg)
	return &cfg
}
