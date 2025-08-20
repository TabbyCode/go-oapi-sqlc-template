// Package server implements HTTP handlers for the user API.
//
//nolint:nilerr // When we use a strict server, we have to return nil error
package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/xurvan/go-oapi-sqlc-template/internal/gen/oapi"
	"github.com/xurvan/go-oapi-sqlc-template/internal/repository"
)

// Server handles HTTP requests for user operations.
type Server struct {
	repo *repository.UserRepository
}

// Ensure that Server implements the StrictServerInterface interface at compile time.
var _ oapi.StrictServerInterface = (*Server)(nil)

// NewServer creates a new Server instance with the given repository.
func NewServer(db *repository.UserRepository) *Server {
	return &Server{
		repo: db,
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

// GetUserById handles GET /users/{id} requests to retrieve a specific user.
func (s *Server) GetUserById(
	ctx context.Context,
	request oapi.GetUserByIdRequestObject,
) (oapi.GetUserByIdResponseObject, error) {
	user, err := s.repo.Get(ctx, request.Id)
	if err != nil {
		return oapi.GetUserById404JSONResponse{
			NotFoundJSONResponse: oapi.NotFoundJSONResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		}, nil
	}

	return oapi.GetUserById200JSONResponse(*user), nil
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
