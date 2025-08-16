package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/xurvan/go-oapi-sqlc-template/internal/config"
	"github.com/xurvan/go-oapi-sqlc-template/internal/database"
	"github.com/xurvan/go-oapi-sqlc-template/internal/gen/oapi"
	"github.com/xurvan/go-oapi-sqlc-template/internal/server"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	srv := server.NewServer(db)
	handler := oapi.HandlerWithOptions(
		srv,
		oapi.GorillaServerOptions{},
	)

	handler = server.SetupMiddlewares(handler)

	httpServer := &http.Server{
		Addr:              cfg.ListenAddress,
		Handler:           handler,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// Channel to listen for errors coming from the listener
	serverErrors := make(chan error, 1)

	// Start the server
	go func() {
		log.Printf("Starting server on %s", cfg.ListenAddress)

		serverErrors <- httpServer.ListenAndServe()
	}()

	// Channel to listen for an interrupt or terminate signal from the OS
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking main and waiting for shutdown or server errors
	select {
	case err := <-serverErrors:
		log.Fatalf("Error starting server: %v", err)

	case <-shutdown:
		log.Println("Shutting down server...")

		// Create a deadline for the shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Gracefully shutdown the server
		err := httpServer.Shutdown(ctx)
		if err != nil {
			log.Printf("Error during server shutdown: %v", err)

			err := httpServer.Close()
			if err != nil {
				log.Printf("Error during server shutdown: %v", err)
			}
		}

		log.Println("Server gracefully stopped")
	}
}
