package config

import (
	"github.com/BurntSushi/toml"
	"github.com/osamikoyo/picall/pkg/loger"
)

type Config struct {
	Addr         string `toml:"address"`
	BucketName   string `toml:"bucket"`
	Secure       bool   `toml:"secure"`
	MinioAddress string `toml:"minio_address"`
}

func New() Config {
	var cfg Config

	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		loger.New().Error().Err(err)
	}

	return cfg
}
