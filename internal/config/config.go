package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

var Config *ConfigModel

func init() {
	initConfig()
}

type ConfigModel struct {
	DB struct {
		Scheme           string `env:"DB_SCHEME" env-default:"postgres"`
		Host             string `env:"DB_HOST" env-default:"127.0.0.1"`
		Port             uint16 `env:"DB_PORT" env-default:"5432"`
		Name             string `env:"DB_NAME" env-default:"milky"`
		Username         string `env:"DB_USERNAME" env-default:"postgres"`
		Password         string `env:"DB_PASSWORD" env-default:"postgres"`
		SSLMode          string `env:"DB_SSL_MODE" env-default:"disable"`
		MaxOpenConns     uint32 `env:"DB_MAX_OPEN_CONNS" env-default:"5"`
		MaxIdleConns     uint32 `env:"DB_MAX_IDLE_CONNS" env-default:"5"`
		MaxConnsLifeTime uint64 `env:"MAX_CONNS_LIFE_TIME" env-default:"600"` //seconds
	}
}

func initConfig() {
	err := cleanenv.ReadEnv(&Config)
	if err != nil {
		log.Fatalf("Read config error: %s", err)
	}
}
