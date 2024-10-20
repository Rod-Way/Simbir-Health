package config

import (
	"account/pkg/db/postgres"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	postgres.Config

	GRPCServerPort int `env:"GRPC_SERVER_PORT" env-default:"9090"`
}

func New() *Config {
	cfg := Config{}
	err := cleanenv.ReadConfig("./configs/local.env", &cfg)
	fmt.Println(err)
	if err != nil {
		return nil
	}
	return &cfg
}
