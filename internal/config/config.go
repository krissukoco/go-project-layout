package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// Server

	Environment     string `envconfig:"ENVIRONMENT" default:"dev"`
	Port            uint   `envconfig:"PORT" default:"31000"`
	ServiceName     string `envconfig:"SERVICE_NAME" default:"go-project-layout"`
	Debug           bool   `envconfig:"SERVICE_NAME" default:"false"`
	GracefulTimeout uint   `envconfig:"GRACEFUL_TIMEOUT" default:"10"` // seconds

	// AUTH

	JwtSecret            string `envconfig:"JWT_SECRET" default:""`
	AccessTokenDuration  uint   `envconfig:"ACCESS_TOKEN_DURATION" default:"120"` // Hours
	RefreshTokenDuration uint   `envconfig:"REFRESH_TOKEN_DURATION" default:"2"`  // Days

	// DB CONFIG

	PostgresHost      string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgresUser      string `envconfig:"POSTGRES_USER" default:"postgres"`
	PostgresPassword  string `envconfig:"POSTGRES_PASSWORD" default:"postgres"`
	PostgresDbname    string `envconfig:"POSTGRES_DBNAME" default:"postgres"`
	PostgresPort      uint   `envconfig:"POSTGRES_PORT" default:"5432"`
	PostgresTimezone  string `envconfig:"POSTGRES_TIMEZONE" default:"Asia/Jakarta"`
	PostgresEnableSsl bool   `envconfig:"POSTGRES_ENABLE_SSL" default:"false"`
}

func Load() (*Config, error) {
	godotenv.Overload()
	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
