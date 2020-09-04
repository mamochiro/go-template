package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

// Configuration ...
type Configuration struct {
	// this section is set env param
	Port    int    `env:"PORT" envDefault:"3000"`
	NodeEnv string `env:"NODE_ENV" envDefault:"local"`

	// DB
	DbDriver   string `env:"DB_DRIVER" envDefault:"postgres"`
	DbHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DbPort     string `env:"DB_PORT" envDefault:"5432"`
	DbUser     string `env:"DB_USER" envDefault:"postgres"`
	DbName     string `env:"DB_NAME" envDefault:"account"`
	DbPassword string `env:"DB_PASSWORD" envDefault:"postgres"`
}

// NewConfiguration ...
func NewConfiguration() Configuration {
	cfg := Configuration{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Errorf("%+v\n", err)
	}
	return cfg
}
