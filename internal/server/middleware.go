package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// SetupMiddlewares configures logging and recovery middlewares for the given handler.
func SetupMiddlewares(handler http.Handler) http.Handler {
	h := handlers.LoggingHandler(os.Stdout, handler)

	h = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(h)

	return h
}
