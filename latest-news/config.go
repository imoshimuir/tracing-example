package main

import (
	"latest-news/telemetry"

	"github.com/ilyakaznacheev/cleanenv"
)

type GlobalConfig struct {
	Port  int    `env:"PORT" env-default:"3333"`
	Host  string `env:"HOST" env-default:"localhost"`
	PostgresConnectionString string `env:"POSTGRES_CONNECTION_STRING" env-default:"host=localhost user=postgres password=password dbname=mydatabase sslmode=disable port=5432"`
	Telemetry telemetry.TracerConfig
}

func LoadConfig() (*GlobalConfig, error) {
	cfg := &GlobalConfig{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}