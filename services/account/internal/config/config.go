package config

import (
	"account/pkg/db/postgres"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	postgres.Config

	HTTPServerPort string `env:"HTTP_SERVER_PORT" env-default:"8080"`
}

func New() *Config {
	cfg := Config{}
	err := cleanenv.ReadConfig("./configs/local.env", &cfg)
	fmt.Println(err)
	if err != nil {
		err := cleanenv.ReadEnv(&cfg)
		fmt.Println(err)
		if err != nil {
			return nil
		}
	}
	return &cfg
}
