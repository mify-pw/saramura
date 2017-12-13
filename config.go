package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/caarlos0/env"
	// autoload .env files
	_ "github.com/joho/godotenv/autoload"
)

var cfg *config

type config struct {
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
	SentryDSN   string `env:"SENTRY"`

	DbHost string `env:"DB_HOST" envDefault:"localhost"`
	DbUser string `env:"DB_USER"`
	DbPass string `env:"DB_PASS"`
	DbName string `env:"DB_NAME"`

	RedisHost string `env:"REDIS_HOST" envDefault:"localhost:6379"`
	RedisPass string `env:"REDIS_PASS"`
	RedisDb   int    `env:"REDIS_DB" envDefault:"0"`
}

func (cfg *config) IsDebug() bool {
	return cfg.Environment == "development"
}

func init() {
	cfg = &config{}
	err := env.Parse(cfg)

	if err != nil {
		log.WithError(err).Fatal("could not read environment")
	}
}
