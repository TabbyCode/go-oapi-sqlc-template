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
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %s", err)
	}

	var cfg Config

	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Failed to load server config: %s", err)
	}

	return &cfg
}
