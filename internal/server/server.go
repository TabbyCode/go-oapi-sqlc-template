// Package server implements HTTP handlers for the user API.
//
//nolint:nilerr,ireturn // The generated code forces us to return interface and nil errors
package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/xurvan/go-oapi-sqlc-template/internal/config"
	"github.com/xurvan/go-oapi-sqlc-template/internal/gen/oapi"
	"github.com/xurvan/go-oapi-sqlc-template/internal/repository"
)

// Server handles HTTP requests for user operations.
type Server struct {
	repo *repository.UserRepository
	app  *fiber.App
	cfg  *config.Config
}

// Ensure that Server implements the StrictServerInterface interface at compile time.
var _ oapi.StrictServerInterface = (*Server)(nil)

// NewServer creates a new Server instance with the given repository.
func NewServer(cfg *config.Config, repo *repository.UserRepository) *Server {
	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	})
	app.Use(logger.New())
	app.Use(recover.New())

	srv := &Server{
		repo: repo,
		app:  app,
		cfg:  cfg,
	}

	strict := oapi.NewStrictHandler(srv, nil)

	oapi.RegisterHandlers(app, strict)

	return srv
}

func (s *Server) Start() {
	// Channel to listen for errors coming from the listener
	serverErrors := make(chan error, 1)

	// Start the server
	go func() {
		log.Printf("Starting server on %s", s.cfg.ListenAddress)

		serverErrors <- s.app.Listen(s.cfg.ListenAddress)
	}()

	// Channel to listen for an interrupt or terminate signal from the OS
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking main and waiting for shutdown or server errors
	select {
	case err := <-serverErrors:
		log.Fatalf("Error starting server: %v", err)

	case <-shutdown:
		err := s.app.ShutdownWithTimeout(s.cfg.ShutdownTimeout)
		if err != nil {
			log.Printf("Error during server shutdown: %v", err)
		}

		log.Println("Server gracefully stopped")
	}
}

// ListUsers handles GET /users requests to retrieve a list of users.
func (s *Server) ListUsers(
	ctx context.Context,
	request oapi.ListUsersRequestObject,
) (oapi.ListUsersResponseObject, error) {
	users, err := s.repo.List(ctx, request.Params)
	if err != nil {
		return oapi.ListUsers500JSONResponse{
			ErrorJSONResponse: oapi.ErrorJSONResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	return oapi.ListUsers200JSONResponse(users), nil
}

// CreateUser handles POST /users requests to create a new user.
func (s *Server) CreateUser(
	ctx context.Context,
	request oapi.CreateUserRequestObject,
) (oapi.CreateUserResponseObject, error) {
	if request.Body == nil {
		return oapi.CreateUser400JSONResponse{
			BadRequestJSONResponse: oapi.BadRequestJSONResponse{
				Code:    http.StatusBadRequest,
				Message: "invalid request body",
			},
		}, nil
	}

	user, err := s.repo.Create(ctx, *request.Body)
	if err != nil {
		return oapi.CreateUser500JSONResponse{
			ErrorJSONResponse: oapi.ErrorJSONResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	return oapi.CreateUser201JSONResponse(*user), nil
}

// GetUserByID handles GET /users/{id} requests to retrieve a specific user.
func (s *Server) GetUserByID(
	ctx context.Context,
	request oapi.GetUserByIDRequestObject,
) (oapi.GetUserByIDResponseObject, error) {
	user, err := s.repo.Get(ctx, request.Id)
	if err != nil {
		return oapi.GetUserByID404JSONResponse{
			NotFoundJSONResponse: oapi.NotFoundJSONResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		}, nil
	}

	return oapi.GetUserByID200JSONResponse(*user), nil
}

// UpdateUser handles PUT /users/{id} requests to update a specific user.
func (s *Server) UpdateUser(
	ctx context.Context,
	request oapi.UpdateUserRequestObject,
) (oapi.UpdateUserResponseObject, error) {
	if request.Body == nil {
		return oapi.UpdateUser400JSONResponse{
			BadRequestJSONResponse: oapi.BadRequestJSONResponse{
				Code:    http.StatusBadRequest,
				Message: "invalid request body",
			},
		}, nil
	}

	user, err := s.repo.Update(ctx, request.Id, *request.Body)
	if err != nil {
		return oapi.UpdateUser500JSONResponse{
			ErrorJSONResponse: oapi.ErrorJSONResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	return oapi.UpdateUser200JSONResponse(*user), nil
}

// DeleteUser handles DELETE /users/{id} requests to delete a specific user.
func (s *Server) DeleteUser(
	ctx context.Context,
	request oapi.DeleteUserRequestObject,
) (oapi.DeleteUserResponseObject, error) {
	err := s.repo.Delete(ctx, request.Id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return oapi.DeleteUser404JSONResponse{
				NotFoundJSONResponse: oapi.NotFoundJSONResponse{
					Code:    http.StatusNotFound,
					Message: "User not found",
				},
			}, nil
		}

		return oapi.DeleteUser500JSONResponse{
			ErrorJSONResponse: oapi.ErrorJSONResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to delete user: " + err.Error(),
			},
		}, nil
	}

	return oapi.DeleteUser204Response{}, nil
}
