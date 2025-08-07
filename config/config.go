package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Address string `envconfig:"ADDRESS" default:"localhost:8080"`
}

func Load() *Config {
	_ = godotenv.Load()

	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Failed to load server config: %s", err)
	}

	return &cfg
}
