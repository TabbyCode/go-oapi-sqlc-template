// Package main provides the HTTP server entry point.
package main

import (
	"log"

	"github.com/xurvan/go-oapi-sqlc-template/internal/config"
	"github.com/xurvan/go-oapi-sqlc-template/internal/repository"
	"github.com/xurvan/go-oapi-sqlc-template/internal/server"
)

func main() {
	cfg := config.Load()

	repo, err := repository.NewUserRepository(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	srv := server.NewServer(cfg, repo)
	srv.Start()
}
