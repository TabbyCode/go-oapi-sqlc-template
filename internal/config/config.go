// Package config handles application configuration loading.
package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config holds all application configuration values.
type Config struct {
	ListenAddress     string        `default:"localhost:8080" envconfig:"LISTEN_ADDRESS"`
	DatabaseURL       string        `                         envconfig:"DATABASE_URL"        required:"true"`
	ReadTimeout       time.Duration `default:"15s"            envconfig:"READ_TIMEOUT"`
	ReadHeaderTimeout time.Duration `default:"5s"             envconfig:"READ_HEADER_TIMEOUT"`
	WriteTimeout      time.Duration `default:"15s"            envconfig:"WRITE_TIMEOUT"`
	IdleTimeout       time.Duration `default:"60s"            envconfig:"IDLE_TIMEOUT"`
	ShutdownTimeout   time.Duration `default:"5s"             envconfig:"SHUTDOWN_TIMEOUT"`
}

// Load reads configuration from environment variables and returns a Config instance.
func Load() *Config {
	_ = godotenv.Load()

	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Failed to load server config: %s", err)
	}

	return &cfg
}
