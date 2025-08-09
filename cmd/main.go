package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/xurvan/go-template/config"
	"github.com/xurvan/go-template/internal/database"
	"github.com/xurvan/go-template/internal/gen/oapi"
	"github.com/xurvan/go-template/internal/server"
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

	httpServer := &http.Server{
		Addr:    cfg.ListenAddress,
		Handler: handler,
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
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Printf("Error during server shutdown: %v", err)
			httpServer.Close()
		}

		log.Println("Server gracefully stopped")
	}
}
