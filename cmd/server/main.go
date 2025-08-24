// Package main provides the HTTP server entry point.
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/xurvan/go-oapi-sqlc-template/internal/config"
	"github.com/xurvan/go-oapi-sqlc-template/internal/gen/oapi"
	"github.com/xurvan/go-oapi-sqlc-template/internal/repository"
	"github.com/xurvan/go-oapi-sqlc-template/internal/server"
)

func main() {
	cfg := config.Load()

	repo, err := repository.NewUserRepository(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	srv := server.NewServer(repo)
	strict := oapi.NewStrictHandler(srv, nil)
	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	})

	app.Use(logger.New())
	app.Use(recover.New())

	oapi.RegisterHandlers(app, strict)

	// Channel to listen for errors coming from the listener
	serverErrors := make(chan error, 1)

	// Start the server
	go func() {
		log.Printf("Starting server on %s", cfg.ListenAddress)

		serverErrors <- app.Listen(cfg.ListenAddress)
	}()

	// Channel to listen for an interrupt or terminate signal from the OS
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking main and waiting for shutdown or server errors
	select {
	case err = <-serverErrors:
		log.Fatalf("Error starting server: %v", err)

	case <-shutdown:
		err = app.ShutdownWithTimeout(cfg.ShutdownTimeout)
		if err != nil {
			log.Printf("Error during server shutdown: %v", err)
		}

		log.Println("Server gracefully stopped")
	}
}
